# *Español* &mdash; Spanish (`es`)

This datasheet is for sps-corpus-3.0-2026-03-09 of the Mozilla Common Voice *Spontaneous Speech* dataset for Spanish [Español - `es`]. The dataset contains 332 clips representing 0.63 hours of recorded speech (0.03 hours validated) from 12 speakers.

## Data splits for modelling

The dataset clips are categorised by transcription status and training-set assignment. The following tables summarise the distribution.

### Audio clips

| Bucket | Clips | % |
| --- | --- | --- |
| Transcribed & Validated | 14 | 4.2% |
| Transcribed & Pending | 3 | 0.9% |
| Not transcribed | 315 | 94.9% |

### Training splits

| Bucket | Clips | % |
| --- | --- | --- |
| Train | 0 | 0.0% |
| Dev | 0 | 0.0% |
| Test | 0 | 0.0% |
| Unassigned | 332 | 100.0% |

Training split coverage: 0 of 14 transcribed & validated clips (0.0%)

## Transcriptions

### Transcription status

| Bucket | Clips | % |
| --- | --- | --- |
| Validated | 14 | 82.4% |
| Pending | 3 | 17.6% |
| Edited | 4 | 23.5% |

### Samples

#### Questions

There follows a randomly selected sample of questions used in the corpus.

1. *Si te sientes aburrido, ¿qué haces?*
2. *¿Cuál es el sitio más bello por donde vives y por qué?*
3. *¿Qué es lo más interesante que has hecho con la Inteligencia Artificial?*
4. *¿Consideras que el cero es un número?*
5. *¿Tienes un libro favorito?*

#### Responses

There follows a randomly selected sample of transcribed responses from the corpus.

1. *¿Qué consejo dan usualmente los mayores a los jóvenes?*
2. *Me gusta la pintura, la música, las artes mundiales*
3. *Lavarse los dientes.*
4. *En algún momento entre la primaria y la secundaria.*
5. *Y por mi región, normalmente, los niños juegan al fútbol.*

### Fields

Each row of a `tsv` file represents a single audio clip, and contains the following information:

- `client_id` - hashed UUID of a given user
- `audio_id` - numeric id for audio file
- `audio_file` - audio file name
- `duration_ms` - duration of audio in milliseconds
- `prompt_id` - numeric id for prompt
- `prompt` - question for user
- `transcription` - transcription of the audio response
- `votes` - number of people that who approved a given transcript
- `age` - age of the speaker[^1]
- `gender` - gender of the speaker[^1]
- `language` - language name
- `split` - for data modelling, which subset of the data does this clip pertain to
- `char_per_sec` - how many characters of transcription per second of audio
- `quality_tags` - some automated assessment of the transcription--audio pair, separated by `|`
  - `transcription-length` - character per second under 3 characters per second
  - `speech-rate` - characters per second over 30 characters per second
  - `short-audio` - audio length under 2 seconds
  - `long-audio` - audio length over 5 minutes

---

[^1]: For a full list of age, gender, and accent options, see the [demographics spec](https://github.com/common-voice/common-voice/blob/main/web/src/stores/demographics.ts). These will only be reported if the speaker opted in to provide that information.

## Get involved

### Community links

- [Common Voice translators on Pontoon](https://pontoon.mozilla.org/es/common-voice/contributors/)
- [Common Voice Communities](https://github.com/common-voice/common-voice/blob/main/docs/COMMUNITIES.md)

### Discussions

- [Common Voice on Matrix](https://chat.mozilla.org/#/room/#common-voice:mozilla.org)
- [Common Voice on Discourse](https://discourse.mozilla.org/t/about-common-voice-readme-first/17218)
- [Common Voice on Discord](https://discord.gg/9QTj9zwn)
- [Common Voice on Telegram](https://t.me/mozilla_common_voice)

### Contribute

- [Contribute questions](https://commonvoice.mozilla.org/spontaneous-speech/beta/question)
- [Validate questions](https://commonvoice.mozilla.org/spontaneous-speech/beta/validate)
- [Answer questions](https://commonvoice.mozilla.org/spontaneous-speech/beta/prompts)
- [Transcribe recordings](https://commonvoice.mozilla.org/spontaneous-speech/beta/transcribe)
- [Validate transcriptions](https://commonvoice.mozilla.org/spontaneous-speech/beta/check-transcript)

## Licence

This dataset is released under the [Creative Commons Zero (CC-0)](https://creativecommons.org/public-domain/cc0/) licence. By downloading this data you agree to not determine the identity of speakers in the dataset.
