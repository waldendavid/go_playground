package openlibrary

import "context"

type Client interface {
	Search(ctx context.Context, req SearchRequest) (SearchResponse, error)
}

type SearchRequest struct {
	Author string
	Title  string
}

type SearchResponse struct {
	NumFound      int  `json:"numFound"`
	Start         int  `json:"start"`
	NumFoundExact bool `json:"numFoundExact"`
	Docs          []struct {
		Key                   string   `json:"key"`
		Type                  string   `json:"type"`
		Seed                  []string `json:"seed"`
		Title                 string   `json:"title"`
		TitleSuggest          string   `json:"title_suggest"`
		HasFulltext           bool     `json:"has_fulltext"`
		EditionCount          int      `json:"edition_count"`
		EditionKey            []string `json:"edition_key"`
		PublishDate           []string `json:"publish_date,omitempty"`
		PublishYear           []int    `json:"publish_year,omitempty"`
		FirstPublishYear      int      `json:"first_publish_year,omitempty"`
		Isbn                  []string `json:"isbn,omitempty"`
		LastModifiedI         int      `json:"last_modified_i"`
		EbookCountI           int      `json:"ebook_count_i"`
		Publisher             []string `json:"publisher,omitempty"`
		PublisherFacet        []string `json:"publisher_facet,omitempty"`
		Version               int64    `json:"_version_"`
		Language              []string `json:"language,omitempty"`
		AuthorKey             []string `json:"author_key,omitempty"`
		AuthorName            []string `json:"author_name,omitempty"`
		Person                []string `json:"person,omitempty"`
		PersonKey             []string `json:"person_key,omitempty"`
		PersonFacet           []string `json:"person_facet,omitempty"`
		AuthorFacet           []string `json:"author_facet,omitempty"`
		PublishPlace          []string `json:"publish_place,omitempty"`
		Contributor           []string `json:"contributor,omitempty"`
		AuthorAlternativeName []string `json:"author_alternative_name,omitempty"`
		IDLibrarything        []string `json:"id_librarything,omitempty"`
		NumberOfPagesMedian   int      `json:"number_of_pages_median,omitempty"`
		IDGoodreads           []string `json:"id_goodreads,omitempty"`
		Lccn                  []string `json:"lccn,omitempty"`
		Lcc                   []string `json:"lcc,omitempty"`
		Subject               []string `json:"subject,omitempty"`
		SubjectFacet          []string `json:"subject_facet,omitempty"`
		LccSort               string   `json:"lcc_sort,omitempty"`
		SubjectKey            []string `json:"subject_key,omitempty"`
		Oclc                  []string `json:"oclc,omitempty"`
		CoverEditionKey       string   `json:"cover_edition_key,omitempty"`
		CoverI                int      `json:"cover_i,omitempty"`
		Ddc                   []string `json:"ddc,omitempty"`
		Ia                    []string `json:"ia,omitempty"`
		PublicScanB           bool     `json:"public_scan_b,omitempty"`
		IaCollectionS         string   `json:"ia_collection_s,omitempty"`
		LendingEditionS       string   `json:"lending_edition_s,omitempty"`
		LendingIdentifierS    string   `json:"lending_identifier_s,omitempty"`
		PrintdisabledS        string   `json:"printdisabled_s,omitempty"`
		FirstSentence         []string `json:"first_sentence,omitempty"`
		Place                 []string `json:"place,omitempty"`
		Time                  []string `json:"time,omitempty"`
		IDAmazon              []string `json:"id_amazon,omitempty"`
		IDDepSitoLegal        []string `json:"id_depósito_legal,omitempty"`
		IDLibrivox            []string `json:"id_librivox,omitempty"`
		IDOverdrive           []string `json:"id_overdrive,omitempty"`
		IDProjectGutenberg    []string `json:"id_project_gutenberg,omitempty"`
		IaLoadedID            []string `json:"ia_loaded_id,omitempty"`
		IaBoxID               []string `json:"ia_box_id,omitempty"`
		PlaceKey              []string `json:"place_key,omitempty"`
		TimeFacet             []string `json:"time_facet,omitempty"`
		PlaceFacet            []string `json:"place_facet,omitempty"`
		DdcSort               string   `json:"ddc_sort,omitempty"`
		TimeKey               []string `json:"time_key,omitempty"`
		IDGoogle              []string `json:"id_google,omitempty"`
	} `json:"docs"`
	NumFound0 int         `json:"num_found"`
	Q         string      `json:"q"`
	Offset    interface{} `json:"offset"`
}
