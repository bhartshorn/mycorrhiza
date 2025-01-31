package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/adrg/xdg"
	"github.com/bouncepaw/mycorrhiza/util"
)

func PopulateFixedUserStorage() {
	contents, err := ioutil.ReadFile(util.FixedCredentialsPath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(contents, &UserStorage.Users)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range UserStorage.Users {
		user.Group = groupFromString(user.GroupString)
	}
	log.Println("Found", len(UserStorage.Users), "fixed users")

	contents, err = ioutil.ReadFile(tokenStoragePath())
	if os.IsNotExist(err) {
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	var tmp map[string]string
	err = json.Unmarshal(contents, &tmp)
	if err != nil {
		log.Fatal(err)
	}
	for token, username := range tmp {
		user := UserStorage.userByName(username)
		UserStorage.Tokens[token] = user
	}
	log.Println("Found", len(tmp), "active sessions")
}

func dumpTokens() {
	tmp := make(map[string]string)
	for token, user := range UserStorage.Tokens {
		tmp[token] = user.Name
	}
	blob, err := json.Marshal(tmp)
	if err != nil {
		log.Println(err)
	} else {
		ioutil.WriteFile(tokenStoragePath(), blob, 0644)
	}
}

// Return path to tokens.json.
func tokenStoragePath() string {
	dir, err := xdg.DataFile("mycorrhiza/tokens.json")
	if err != nil {
		// Yes, it is unix-only, but function above rarely fails, so this block is probably never reached.
		path := "/home/" + os.Getenv("HOME") + "/.local/share/mycorrhiza/tokens.json"
		os.MkdirAll(path, 0777)
		return path
	}
	return dir
}
