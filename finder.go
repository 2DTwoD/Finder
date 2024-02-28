package main

import (
	"Finder/globals"
	"Finder/pathEntry"
	"Finder/utils"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const startPath string = "D:/Virtual Machines/" //"./"
const startFilter string = "wtf"

func main() {
	log.Println("Start")
	//text := "MRVN\u0001   CMOStimA\u0004   PF  CMOStimB\b   PF      CMOS    Ђ             &\u0002 Ђ  @ђр4\u0003Ђ\u0002 |\u0001 ¬Џh¬\u000F\u0004\u0004\u0004\u0004\u0002\u0006р;\u0004я\u0001  Ђ\b\u001B | АЖ‘  шаЂ,А         €\u000FрG4Ц'          \n      а±ђ      \u0098#O                      CMOS440 Ђ         \u0001   ¦\u0002 Ђ   ’ €\u0003Ђ\u0002 ь  c}Dc} Т\u0006\n   Ф\u0003”\u0003   \u0006п ь ИР\u0011\u0001   А™\u0013\u001E    @і\u0001 \u0018Њ€ \u0004       \u0001ю \u001Aл\v          \u0014      \u0018 \u0016   @N›яяд'DЪP       CMOS440v\u0004   \u0004   CMOS440xЂ   hЂ0\u0018 ѓ\u0001Ђ‡aЎ      pЕэя_! Т†тИ\u0003                                                                                                   ESCD        NAPI\u001F ]BBSI\u001F ћ\u0001ESCDЅ\u0001К\u0011ELOG‡\u0013 \u0005 \u0002\u0001                                                                                                                                        B  Л\u0004\u0011          CD\u0001123456789ABCDEC                                                                                                         Иљ\u0002          B                                                                                                                        qџК\u0011ACFG\u0001\u00022   `\u0001      P@  \n \u0001 \u0012 \n юяЂ  Ђ B \u0001 \u0002Ѓ\n   Ђ\u0002Ђ\nЂ\u000E ` Ђ\n \u0010 ьяЂ\nр\u000F\u0004ьяЂ\nа\u000F\bьяЂ\nР\u000F\fьяЂ\nА\u000F\u0010ьяЂ\n°\u000F\u0014ья \n \u000F\u0018ья\u000E \u0001 \u0018\u0004\fЏ  ЋЃ \u001FА \v \u0001 \u0014\u0002 Ѓ  \u0001  \b \u0001 \u0014  \u0003@ \b \u0001 \u0014\b \u0001p \v \u0001 \u0014\u0001 Ђ`  d \b \u0001 \u0014\n \u000Fр \u0006 \u0001 \u0010 a \n \u0001 \u0002\u0001\nА\n \u0010 \u0006 \u0001 \u0010\aш\f? \u0001 \u0010ЃР\u0004џ \u0010џ \u0010џ@\u0010€\u0010 Ђ\u001F Ѓ$ Ѓ( Ѓ, Ѓ0 Ѓ4 Ѓ8 Ѓ< ‚P …r Џђ Ѓ¤ ЃЁ Ѓ¬ \n° \u0011 \u0001 \u0002Ђ\n Аю@  \n аю\u0004 \u0003 \u0001 Ђ\n \u0001 \u0002 \nР\f \f \u0005 \u0001 \u0004\f \b \u0001 \u0014D \aш\u0003\b \u0001 \u0014C \aш\u0002\b \u0001 \u0014\a \ax\u0003\n \u0001 \u001C\u0006 \u0002 …р\u0003 ч\u0003    R \u0010     `@  \" \u0001 ’\u0098\n  рья\u0098\nряуья\u0098\nаяч\b \u001B\n Ђю \b\u001Fаь\u001C \u0001 А\u0018ACFG \u0002\u0004           ђq†Ђ      o \u0011     `@  \u0016 \u0001 ”Џ \u000E Ђv\u0003‡p\u0001Ђц\u0003‡р\u0001\u000F`\u0010\u0003 \u0001 Ђ\u0012 \u0001 –\u001B\nаїю\b i џЂ\u0010\u001F \u0010, \u0001 А(ACFG \u0002\u0004          9\u0011q†Ђ   ;\u0013q†Ђ   ?@\a­\u0015      b \u0012     `@  2 \u0001 –љ\n \f   ™\n \n Ђ ›\n  ю  ›\n  иья›\nрялья\u001B\nаяп\b i \u000Fp\u0010\u001C \u0001 А\u0018ACFG \u0002\u0004          x\u0005\u0004­\u0015      3 \u0013     `@  \u0003 \u0001 Ђ\u001C \u0001 А\u0018ACFG \u0002\u0004          €ђ\a­\u0015      8 \u0014     `@  \b \u0001 ”k \u001F@ \u001C \u0001 А\u0018ACFG \u0002\u0004         \u0002 t\a­\u0015      P \u0015     `@    \u0001 –љ\nЂ\f \u0004 ›\n яэ@ \u001B\n \\эЂ j џ  \u001F  \u001C \u0001 А\u0018ACFG \u0002\u0004         \u0002\b\u000F\u0010†Ђ      < \u0016     `@  \f \u0001 †\u001B\nАюэ\u0010 i \u001C \u0001 А\u0018ACFG \u0002\u0004         \u0002\u0010w\u0019­\u0015      < \u0017     `@  \f \u0001 †\u001B\n°^э\u0004 e \u001C \u0001 А\u0018ACFG \u0002\u0004         \u0002\u0018p\a­\u0015      C \u0018     `@  \u0013 \u0001 †љ\nђ\f \b \u001B\n ^э\u0004 j \u001C \u0001 А\u0018ACFG \u0002\u0004         \u0002(а\a­\u0015      b \u0019     `@  2 \u0001 –љ\n°\f \b ›\n Oэ@ \u001B\nАNэ\u0010 k џ @џ @џ@@џ`@џЂ@џ @џА@\u001Fа@\u001C \u0001 А\u0018ACFG \u0002\u0004         \u0003 T  \u0010      Ѓ \u001A     `@  \u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 ЂL \u0001 АHACFG \u0002\u0004          © \a­\u0015   Є \a­\u0015   « \a­\u0015   ¬ \a­\u0015   ­ \a­\u0015   ® \a­\u0015   Ї \a­\u0015      Ћ \u001B     `@  \u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 ЂT \u0001 АPACFG \u0002\u0004          ° \a­\u0015   ± \a­\u0015   І \a­\u0015   і \a­\u0015   ґ \a­\u0015   µ \a­\u0015   ¶ \a­\u0015   · \a­\u0015      Ћ \u001C     `@  \u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 ЂT \u0001 АPACFG \u0002\u0004          ё \a­\u0015   № \a­\u0015   є \a­\u0015   » \a­\u0015   ј \a­\u0015   Ѕ \a­\u0015   ѕ \a­\u0015   ї \a­\u0015      Ћ \u001D     `@  \u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 Ђ\u0003 \u0001 ЂT \u0001 АPACFG \u0002\u0004          А \a­\u0015   Б \a­\u0015   В \a­\u0015   Г \a­\u0015   Д \a­\u0015   Е \a­\u0015   Ж \a­\u0015   З \a­\u0015      R \u001E     `@  @ \u0001 А<ACFG \u0002@           ђq†Ђ  `\u0010  яь    Pыяяїю  Ае    яяџю            R \u001F     `@  @ \u0001 А<ACFG \u0002@          \b‘q†Ђ\u0001\u0001 ряяя\u000F    ряяя\u000F   ря    яя\u000F             R       `@  @ \u0001 А<ACFG \u0002@          €ђ\a­\u0015\u0002\u0002    я?    Pэяяяэ  °з    яяяз            R !     `@  @ \u0001 А<ACFG \u0002@          Ё \a­\u0015\u0003\u0003 @  яO    @эяяOэ  ря    яя\u000F             R \"     `@  @ \u0001 А<ACFG \u0002@          © \a­\u0015\u0004\u0004 Ђ  яЏ     эяя\u000Fэ  pз    яя\u007Fз            R #     `@  @ \u0001 А<ACFG \u0002@          Є \a­\u0015\u0005\u0005 А  яП    АьяяПь  0з    яя?з            R $     `@  @ \u0001 А<ACFG \u0002@          « \a­\u0015\u0006\u0006 ряяя\u000F    ЂьяяЏь  рж    яяяж            R %     `@  @ \u0001 А<ACFG \u0002@          ¬ \a­\u0015\a\a ряяя\u000F    @ьяяOь  °ж    яяїж            R &     `@  @ \u0001 А<ACFG \u0002@          ­ \a­\u0015\b\b ряяя\u000F     ьяя\u000Fь  pж    яя\u007Fж            R '     `@  @ \u0001 А<ACFG \u0002@          ® \a­\u0015\t\t ряяя\u000F    АыяяПы  0ж    яя?ж            R (     `@  @ \u0001 А<ACFG \u0002@          Ї \a­\u0015\n\n ряяя\u000F    ЂыяяЏы  ре    яяяе            R )     `@  @ \u0001 А<ACFG \u0002@          ° \a­\u0015\v\v P  я_    0эяя?э   з    яяЇз            R *     `@  @ \u0001 А<ACFG \u0002@          ± \a­\u0015\f\f ђ  яџ    рьяяяь  `з    яяoз            R +     `@  @ \u0001 А<ACFG \u0002@          І \a­\u0015\n\n Р  яЯ    °ьяяїь   з    яя/з            R ,     `@  @ \u0001 А<ACFG \u0002@          і \a­\u0015\u000E\u000E ряяя\u000F    pьяя\u007Fь  аж    яяпж            R -     `@  @ \u0001 А<ACFG \u0002@          ґ \a­\u0015\u000F\u000F ряяя\u000F    0ьяя?ь   ж    яяЇж            R .     `@  @ \u0001 А<ACFG \u0002@          µ \a­\u0015\u0010\u0010 ряяя\u000F    рыяяяы  `ж    яяoж            R /     `@  @ \u0001 А<ACFG \u0002@          ¶ \a­\u0015\u0011\u0011 ряяя\u000F    °ыяяїы   ж    яя/ж            R 0     `@  @ \u0001 А<ACFG \u0002@          · \a­\u0015\u0012\u0012 ряяя\u000F    pыяя\u007Fы  ае    яяпе            R 1     `@  @ \u0001 А<ACFG \u0002@          ё \a­\u0015\u0013\u0013 `  яo     эяя/э  ђз    яяџз            R 2     `@  @ \u0001 А<ACFG \u0002@          № \a­\u0015\u0014\u0014    яЇ    аьяяпь  Pз    яя_з            R 3     `@  @ \u0001 А<ACFG \u0002@          є \a­\u0015\u0015\u0015 а  яп     ьяяЇь  \u0010з    яя\u001Fз            R 4     `@  @ \u0001 А<ACFG \u0002@          » \a­\u0015\u0016\u0016 ряяя\u000F    `ьяяoь  Рж    яяЯж            R 5     `@  @ \u0001 А<ACFG \u0002@          ј \a­\u0015\u0017\u0017 ряяя\u000F     ьяя/ь  ђж    яяџж            R 6     `@  @ \u0001 А<ACFG \u0002@          Ѕ \a­\u0015\u0018\u0018 ряяя\u000F    аыяяпы  Pж    яя_ж            R 7     `@  @ \u0001 А<ACFG \u0002@          ѕ \a­\u0015\u0019\u0019 ряяя\u000F     ыяяЇы  \u0010ж    яя\u001Fж            R 8     `@  @ \u0001 А<ACFG \u0002@          ї \a­\u0015\u001A\u001A ряяя\u000F    `ыяяoы  Ре    яяЯе            R 9     `@  @ \u0001 А<ACFG \u0002@          А \a­\u0015\u001B\u001B p  я\u007F    \u0010эяя\u001Fэ  Ђз    яяЏз            R :     `@  @ \u0001 А<ACFG \u0002@          Б \a­\u0015\u001C\u001C °  яї    РьяяЯь  @з    яяOз            R ;     `@  @ \u0001 А<ACFG \u0002@          В \a­\u0015\u001D\u001D ряяя\u000F    ђьяяџь   з    яя\u000Fз            R <     `@  @ \u0001 А<ACFG \u0002@          Г \a­\u0015\u001E\u001E ряяя\u000F    Pьяя_ь  Аж    яяПж            R =     `@  @ \u0001 А<ACFG \u0002@          Д \a­\u0015\u001F\u001F ряяя\u000F    \u0010ьяя\u001Fь  Ђж    яяЏж            R >     `@  @ \u0001 А<ACFG \u0002@          Е \a­\u0015   ряяя\u000F    РыяяЯы  @ж    яяOж            R ?     `@  @ \u0001 А<ACFG \u0002@          Ж \a­\u0015!! ряяя\u000F    ђыяяџы   ж    яя\u000Fж            R @     `@  @ \u0001 А<ACFG \u0002@          З \a­\u0015\"\" ряяя\u000F    Pыяя_ы  Ае    яяПе            ї:     `\u0001} |\u0001}   \u0001яяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяяя                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        л"
	//println(text, strings.ToLower(text))
	//println(len(text), len(strings.ToLower(text)))
	//return
	globals.SetCurrentFileName(filepath.Base(os.Args[0]))
	if startFilter == "" {
		globals.SetFilter(strings.ToLower(strings.TrimSuffix(globals.GetCurrentFileName(), filepath.Ext(globals.GetCurrentFileName()))))
	} else {
		globals.SetFilter(strings.ToLower(startFilter))
	}
	globals.SetResultFileName(utils.GetResultFileName())
	resultFile, err := os.Create(globals.GetResultFileName())
	if err != nil {
		log.Fatal(err)
	}
	defer func(resultFile *os.File) {
		err := resultFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resultFile)

	utils.WriteLine(resultFile, utils.GetHeaderLine(utils.GetAbsolutePath(startPath)))

	var dirEntriesWithPath = make([]*pathEntry.DirEntryWithPath, 0)

	dirEntries, err := os.ReadDir(startPath)
	if err != nil {
		log.Fatal(err)
	}
	dirEntriesWithPath = getEntryPathSlice(dirEntries, startPath)

	var i int

	log.Println("Creating folders tree...")

	for {
		if i >= len(dirEntriesWithPath) {
			break
		}
		dirEntryWithPath := dirEntriesWithPath[i]
		i++
		if dirEntryWithPath.IsDir() {
			dirEntryWithPath.AppendPath()
			moreDirEntries, _ := os.ReadDir(dirEntryWithPath.Path())
			dirEntriesWithPath = append(dirEntriesWithPath, getEntryPathSlice(moreDirEntries, dirEntryWithPath.Path())...)
		}
	}

	log.Println("Search for matches...")

	waitChan := make(chan bool, runtime.NumCPU())
	for _, dirEntryWithPath := range dirEntriesWithPath {
		waitChan <- true
		go utils.SearchScript(dirEntryWithPath, resultFile, waitChan)
	}

	for len(waitChan) != 0 {
	}

	log.Println("Finish")

	utils.WriteLine(resultFile, utils.GetEndLine())
}

func getEntryPathSlice(dirEntrySlice []os.DirEntry, path string) []*pathEntry.DirEntryWithPath {
	result := make([]*pathEntry.DirEntryWithPath, 0)
	for _, dirEntry := range dirEntrySlice {
		result = append(result, pathEntry.New(&dirEntry, path))
	}
	return result
}
