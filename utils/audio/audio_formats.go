package audio

// AudioFormat defines the type of the audio
type AudioFormat string

const (
	// WAV : .wav
	WAV AudioFormat = "wav"
	// MP3 : mp3
	MP3 AudioFormat = "mp3"
	// AAC : m4a
	AAC AudioFormat = "m4a"
	// FLAC : flac
	FLAC AudioFormat = "flac"
)

// GetAllFormats gets all the types of audios to transform
func GetAllFormats() []AudioFormat {
	return []AudioFormat{
		WAV,
		MP3,
		AAC,
		FLAC,
	}
}

// ToString converts AudioFormat to string
func (audioType AudioFormat) ToString() string {
	switch audioType {
	case WAV:
		return "wav"
	case MP3:
		return "mp3"
	case AAC:
		return "m4a"
	case FLAC:
		return "flac"
	default:
		return "wav"
	}
}

// ToAudioType converts string to AudioFormat
func ToAudioType(
	audioType string,
) AudioFormat {
	switch audioType {
	case "wav":
		return WAV
	case "mp3":
		return MP3
	case "m4a":
		return AAC
	case "flac":
		return FLAC
	default:
		return WAV
	}
}
