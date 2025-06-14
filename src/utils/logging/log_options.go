package logging

type LogOptions struct {
	Show           bool
	Msg            string
	Tag            string
	LogOptionsType LogOptionsType
}

func (super LogOptions) UpdateShow(show bool) LogOptions {
	return LogOptions{
		Show:           show,
		Msg:            super.Msg,
		Tag:            super.Tag,
		LogOptionsType: super.LogOptionsType,
	}
}

func (super LogOptions) UpdateMsg(msg string) LogOptions {
	return LogOptions{
		Show:           super.Show,
		Msg:            msg,
		Tag:            super.Tag,
		LogOptionsType: super.LogOptionsType,
	}
}

func (super LogOptions) UpdateTag(tag string) LogOptions {
	return LogOptions{
		Show:           super.Show,
		Msg:            super.Msg,
		Tag:            tag,
		LogOptionsType: super.LogOptionsType,
	}
}

func (super LogOptions) UpdateType(logType LogOptionsType) LogOptions {
	return LogOptions{
		Show:           super.Show,
		Msg:            super.Msg,
		Tag:            super.Tag,
		LogOptionsType: logType,
	}
}
