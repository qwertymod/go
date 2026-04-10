package cloud


type cloud struct {
	url string
}


func NewCloud(url string) *cloud{
	return &cloud{
		url : url,
	}
}

func (*cloud) Write([]byte) ()

func (*cloud) Read() ([]byte, error) {
	return []byte{}, nil
}