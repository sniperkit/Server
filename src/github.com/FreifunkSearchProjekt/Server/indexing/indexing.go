package indexing

import (
	"encoding/base64"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/analysis/analyzer/custom"
	"github.com/blevesearch/bleve/analysis/token/lowercase"
	"github.com/blevesearch/bleve/analysis/token/ngram"
	"github.com/blevesearch/bleve/analysis/tokenizer/single"
	"github.com/blevesearch/bleve/analysis/tokenizer/unicode"
	"github.com/blevesearch/bleve/mapping"
	"github.com/blevesearch/bleve/search/highlight/format/html"
	"log"
	"os"
	"sync"
)

type Indexer struct {
	idxs map[string]bleve.Index
	sync.RWMutex
}

func (i *Indexer) getIndex(id string) (idx bleve.Index) {
	i.RLock()
	idx = i.idxs[id]
	i.RUnlock()

	if idx != nil {
		return
	}

	i.Lock()
	idx, _ = Bleve(base64.URLEncoding.EncodeToString([]byte(id)))
	i.idxs[id] = idx
	i.Unlock()
	return
}

func (i *Indexer) AddBasicWebpage(ID, CommunityID string, wp WebpageBasic) {
	wp.Index(ID, i.getIndex(CommunityID))
}

func (i *Indexer) AddBasicFeed(ID, CommunityID string, fb FeedBasic) {
	fb.Index(ID, i.getIndex(CommunityID))
}

func (i *Indexer) GetFields(CommunityID string) ([]string, error) {
	return i.getIndex(CommunityID).Fields()
}

func (i *Indexer) Query(id, query string) (*bleve.SearchResult, error) {
	//searchRequest := bleve.NewSearchRequest(bleve.NewMatchQuery(query))
	//searchRequest := bleve.NewSearchRequest(bleve.NewFuzzyQuery(query))
	//searchRequest := bleve.NewSearchRequest(bleve.NewQueryStringQuery(query))
	searchTerm := bleve.NewQueryStringQuery(query)
	searchRequest := bleve.NewSearchRequest(searchTerm)
	searchRequest.Fields = make([]string, 5)
	searchRequest.Fields[0] = "URL"
	searchRequest.Fields[1] = "Host"
	searchRequest.Fields[2] = "Path"
	searchRequest.Fields[3] = "Title"
	searchRequest.Fields[4] = "Description"
	searchRequest.Highlight = bleve.NewHighlightWithStyle(html.Name)
	return i.getIndex(id).Search(searchRequest)
}

func NewIndexer() Indexer {
	return Indexer{
		idxs: make(map[string]bleve.Index),
	}
}

//var bleveIdx bleve.Index
//var bleveIdxMap = make(map[string]bleve.Index)

// Bleve connect or create the index persistence
func Bleve(indexPath string) (bleve.Index, error) {

	//if bleveIdx, exists := bleveIdxMap[indexPath]; exists {
	//	return bleveIdx, nil
	//}

	// try to open de persistence file...
	bleveIdx, err := bleve.Open("bleve/" + indexPath)

	// if doesn't exists or something goes wrong...
	if err != nil {
		// create a new mapping file and create a new index
		var newMapping mapping.IndexMapping
		newMapping, err = buildIndexMapping()
		if err != nil {
			return nil, err
		}
		bleveIdx, err = bleve.New("bleve/"+indexPath, newMapping)
		if err != nil {
			return nil, err
		}
	}

	//if err == nil {
	//	bleveIdxMap[indexPath] = bleveIdx
	//}
	return bleveIdx, err
}

func OpenIndex(databasePath string) bleve.Index {
	index, err := bleve.Open(databasePath)

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	return index
}

func CreateIndex(databasePath string) bleve.Index {
	mapping := bleve.NewIndexMapping()
	mapping = addCustomAnalyzers(mapping)
	mapping = createEventMapping(mapping)

	index, err := bleve.New(databasePath, mapping)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	return index
}

func createEventMapping(indexMapping *mapping.IndexMappingImpl) *mapping.IndexMappingImpl {
	eventMapping := bleve.NewDocumentMapping()

	eventIDMapping := bleve.NewTextFieldMapping()
	eventIDMapping.IncludeInAll = false
	eventMapping.AddFieldMappingsAt("event_id", eventIDMapping)

	senderMapping := bleve.NewTextFieldMapping()
	senderMapping.IncludeInAll = false
	eventMapping.AddFieldMappingsAt("sender", senderMapping)

	roomIDMapping := bleve.NewTextFieldMapping()
	roomIDMapping.IncludeInAll = false
	eventMapping.AddFieldMappingsAt("room_id", roomIDMapping)

	contentMapping := bleve.NewTextFieldMapping()
	contentMapping.IncludeInAll = false
	eventMapping.AddFieldMappingsAt("content", contentMapping)

	indexMapping.AddDocumentMapping("event", eventMapping)

	return indexMapping
}

func addCustomTokenFilter(indexMapping *mapping.IndexMappingImpl) *mapping.IndexMappingImpl {
	err := indexMapping.AddCustomTokenFilter("bigram_tokenfilter", map[string]interface{}{
		"type": ngram.Name,
		//"side": ngram.FRONT,
		"min": 3.0,
		"max": 25.0,
	})

	if err != nil {
		log.Fatal(err)
	}

	return indexMapping
}

func addCustomAnalyzers(indexMapping *mapping.IndexMappingImpl) *mapping.IndexMappingImpl {
	indexMapping = addCustomTokenFilter(indexMapping)

	err := indexMapping.AddCustomAnalyzer("not_analyzed", map[string]interface{}{
		"type":      custom.Name,
		"tokenizer": single.Name,
	})

	if err != nil {
		log.Fatal(err)
	}

	err = indexMapping.AddCustomAnalyzer("fulltext_ngram", map[string]interface{}{
		"type":      custom.Name,
		"tokenizer": unicode.Name,
		"token_filters": []string{
			lowercase.Name,
			"bigram_tokenfilter",
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	return indexMapping
}

func Execute() {

}
