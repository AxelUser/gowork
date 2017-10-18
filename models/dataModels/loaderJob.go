package dataModels

//LoaderJob is to store data to load vacancies per skill
type LoaderJob struct {
	URL   string
	Alias string
}

//NewLoaderJob creates LoaderJob
func NewLoaderJob(alias string, url string) LoaderJob {
	return LoaderJob{
		Alias: alias,
		URL:   url,
	}
}
