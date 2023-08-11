package config

type UploadConf struct {
	VideoLocalPath string `json:"video_local_path" mapstructure:"video_local_path"`
	CoverLocalPath string `json:"cover_local_path" mapstructure:"cover_local_path"`
	UrlPath        string `json:"url_path" mapstructure:"url_path"`
}
