# Anki Chinese CSV convertor (ACCC)

## Goal

Create an cli application to update exported anki decks into csv files to help with Mandarin language learning. 

```
$ tasks
```

## Requirements

Should be able to perform crud operations via a cli on a data file of tasks. The operations should be as follows:

```
$ acc convert

Where is the folder located at? (default /Users/michaelomh/Downloads)
_

What is the file name? 
_

chinese  english   pinyin
挑战	    challenge Tiǎozhàn
```

### Convert

The convert method should get the folder location and file name and convert it into "chinese, english, pinyin".

## Notable Packages Used

- `encoding/csv` for writing out as a csv file
- `strconv` for turning types into strings and visa versa
- `text/tabwriter` for writing out tab aligned output
- `os` for opening and reading files
- `github.com/spf13/cobra` for the command line interface
- [Optional] `github.com/charmbracelet/huh` if you prefer a TUI interface instead.
 
## Custom Resources

### Example Data File

Additionally, an example txt file would look something like this.

```
#separator:tab
#html:true
#tags column:3
挑战	challenge<br><br>Tiǎozhàn	
点赞	like (social media)<br>Diǎn zàn<br><br>点 - click<br>赞 - praise	
代替	Replace / Substitute<br>Dàitì	
商场	Shopping Mall<br>Shāngchǎng	
口音	accent<br>Kǒuyīn	
电子邮件	email<br>Diànzǐ yóujiàn<br>	
注意	Take note of<br>Zhùyì	
联络 / 联系	Contact<br>Liánluò / Liánxì	
主人公	Protagonist<br>Zhǔréngōng	
终于	at last / finally<br>Zhōngyú	
化妆	put on makeup<br>Huàzhuāng	
吃惊	Shocked / Startled<br>Chījīng	
一切顺利	wishing&nbsp;all the best<br>Yīqiè shùnlì	
技术	Technology<br>Jìshù	
懒	lazy<br>Lǎn	
戴上	put on (usually on the head)<br>Dài shàng	
澳大利亚	Australia<br>Àodàlìyǎ	
缺点	disadvantage / shortcomings<br>Quēdiǎn	
危险	danger<br>Wéixiǎn	
控制	control<br>Kòngzhì	
熬	simmer (soup) / endure<br>Áo	
气氛	vibe / atmosphere<br>Qìfēn	
```

## Run instructions
`make build` - create a file in the `/usr/local/bin/` folder called task. This would allow you to run `tasks <action>`

`make clean` - removes the tasks file in the `/usr/local/bin/`
