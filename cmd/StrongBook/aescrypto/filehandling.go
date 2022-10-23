package aescrypto

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	Random "math/rand"
	"os"
	"strings"
	"time"
)

var WORD_LIST = [...]string{"the", "of", "to", "and", "a", "in", "is", "it", "you", "that", "he", "was", "for", "on", "are", "with", "as", "I", "his", "they", "be", "at", "one", "have", "this", "from", "or", "had", "by", "hot", "but", "some", "what", "there", "we", "can", "out", "other", "were", "all", "your", "when", "up", "use", "word", "how", "said", "an", "each", "she", "which", "do", "their", "time", "if", "will", "way", "about", "many", "then", "them", "would", "write", "like", "so", "these", "her", "long", "make", "thing", "see", "him", "two", "has", "look", "more", "day", "could", "go", "come", "did", "my", "sound", "no", "most", "number", "who", "over", "know", "water", "than", "call", "first", "people", "may", "down", "side", "been", "now", "find", "any", "new", "work", "part", "take", "get", "place", "made", "live", "where", "after", "back", "little", "only", "round", "man", "year", "came", "show", "every", "good", "me", "give", "our", "under", "name", "very", "through", "just", "form", "much", "great", "think", "say", "help", "low", "line", "before", "turn", "cause", "same", "mean", "differ", "move", "right", "boy", "old", "too", "does", "tell", "sentence", "set", "three", "want", "air", "well", "also", "play", "small", "end", "put", "home", "read", "hand", "port", "large", "spell", "add", "even", "land", "here", "must", "big", "high", "such", "follow", "act", "why", "ask", "men", "change", "went", "light", "kind", "off", "need", "house", "picture", "try", "us", "again", "animal", "point", "mother", "world", "near", "build", "self", "earth", "father", "head", "stand", "own", "page", "should", "country", "found", "answer", "school", "grow", "study", "still", "learn", "plant", "cover", "food", "sun", "four", "thought", "let", "keep", "eye", "never", "last", "door", "between", "city", "tree", "cross", "since", "hard", "start", "might", "story", "saw", "far", "sea", "draw", "left", "late", "run", "don't", "while", "press", "close", "night", "real", "life", "few", "stop", "open", "seem", "together", "next", "white", "children", "begin", "got", "walk", "example", "ease", "paper", "often", "always", "music", "those", "both", "mark", "book", "letter", "until", "mile", "river", "car", "feet", "care", "second", "group", "carry", "took", "rain", "eat", "room", "friend", "began", "idea", "fish", "mountain", "north", "once", "base", "hear", "horse", "cut", "sure", "watch", "color", "face", "wood", "main", "enough", "plain", "girl", "usual", "young", "ready", "above", "ever", "red", "list", "though", "feel", "talk", "bird", "soon", "body", "dog", "family", "direct", "pose", "leave", "song", "measure", "state", "product", "black", "short", "numeral", "class", "wind", "question", "happen", "complete", "ship", "area", "half", "rock", "order", "fire", "south", "problem", "piece", "told", "knew", "pass", "farm", "top", "whole", "king", "size", "heard", "best", "hour", "better", "TRUE", "during", "hundred", "am", "remember", "step", "early", "hold", "west", "ground", "interest", "reach", "fast", "five", "sing", "listen", "six", "table", "travel", "less", "morning", "ten", "simple", "several", "vowel", "toward", "war", "lay", "against", "pattern", "slow", "center", "love", "person", "money", "serve", "appear", "road", "map", "science", "rule", "govern", "pull", "cold", "notice", "voice", "fall", "power", "town", "fine", "certain", "fly", "unit", "lead", "cry", "dark", "machine", "note", "wait", "plan", "figure", "star", "box", "noun", "field", "rest", "correct", "able", "pound", "done", "beauty", "drive", "stood", "contain", "front", "teach", "week", "final", "gave", "green", "oh", "quick", "develop", "sleep", "warm", "free", "minute", "strong", "special", "mind", "behind", "clear", "tail", "produce", "fact", "street", "inch", "lot", "nothing", "course", "stay", "wheel", "full", "force", "blue", "object", "decide", "surface", "deep", "moon", "island", "foot", "yet", "busy", "test", "record", "boat", "common", "gold", "possible", "plane", "age", "dry", "wonder", "laugh", "thousand", "ago", "ran", "check", "game", "shape", "yes", "hot", "miss", "brought", "heat", "snow", "bed", "bring", "sit", "perhaps", "fill", "east", "weight", "language", "among"}
var HEADER_LIST = [...]string{"Captin Davey Smack", "Bulit", "Nicky G", "JP"}

func EncryptFile(password string, inputFile string, outputFile string) {
	inputFileContents, inputFileErr := ioutil.ReadFile(inputFile)
	Random.Seed(time.Now().UnixNano())

	if inputFileErr != nil {
		log.Fatal(inputFileErr)
	}
	var salt = generateRandomSalt(32 - len(password))
	var hashedPassword = hashPassword(password, salt)

	iv := make([]byte, 16)
	rand.Read(iv)
	AESEncrytped := EncryptAES(string(inputFileContents), hashedPassword, iv, aes.BlockSize)
	convertedMsg := convertMessageEncrypt(AESEncrytped)
	header := HEADER_LIST[Random.Intn(len(HEADER_LIST)-0)+0] + "\n\n"
	encryptedMsgContents := header + convertedMsg
	fmt.Println("Password Bytes Length: ", len(hashedPassword))
	fmt.Println("Salt Length:", len(salt))
	fmt.Println("Header:", header)
	fmt.Println("Msg:", encryptedMsgContents)
	writeFile(encryptedMsgContents, outputFile)
}

func convertMessageEncrypt(encryptedArr []byte) string {
	encryptedMsg := ""
	wordCounter := 0

	for msgByte := range encryptedArr {
		randPeriod := Random.Intn(2)
		if wordCounter == 0 {
			encryptedMsg += strings.Title(WORD_LIST[msgByte]) + " "
		} else if (randPeriod == 1) && (wordCounter >= 10) && (wordCounter <= 15) {
			encryptedMsg += WORD_LIST[msgByte] + "." + " "
			wordCounter = -1
		} else {
			encryptedMsg += WORD_LIST[msgByte] + " "
		}
		wordCounter++
	}
	return encryptedMsg
}
func writeFile(msg string, outfile string) {
	f, err := os.Create(outfile)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(msg)

	if err2 != nil {
		log.Fatal(err2)
	}
}
