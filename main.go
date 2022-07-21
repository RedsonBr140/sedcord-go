package main

import (
  "errors"
  "fmt"
  "io"
  "log"
  "os"
  "os/exec"
  "os/signal"
  "strings"
  "syscall"

  "github.com/bwmarrin/discordgo"
  "github.com/joho/godotenv"
)


func main(){
  if err := godotenv.Load(); err != nil {
    log.Fatal("Error loading the .env file,", err)
  }

  token := os.Getenv("TOKEN")

  dg, err := discordgo.New("Bot " + token)
  if err != nil {
    fmt.Println("Error creation Discord section,", err)
    return
  }
  
  dg.AddHandler(messageCreate)

  if err := dg.Open(); err != nil {
    fmt.Println("Chore:", err)
    return
  }

  fmt.Println("Running, press CTRL + C to exit.")
  sc := make(chan os.Signal, 1)
  signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
  <-sc

  dg.Close()
}

func sed(sedargs string, inputmsg string) (string, error) {
  cmd := exec.Command("sed", "--sandbox", sedargs)
  stdin, err := cmd.StdinPipe()
  if err != nil {
    return "", errors.New("Could not connect to Stdin")
  }
  
  go func() {
    defer stdin.Close()
    io.WriteString(stdin, inputmsg)
  }()

  out, err := cmd.CombinedOutput()
  if err != nil {
    return "", errors.New("Could not get Stdout")
  }

  return string(out), nil
}


func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){
  if m.Author.ID == s.State.User.ID {
    return
  }

  // If message has prefix and is a reply to another message.
  if strings.Contains(m.Content, "!s") && m.Message.Type == 19 {
    contentSlice := strings.Split(m.Content, " ")
    contentSlice = append(contentSlice[:0], contentSlice[1:]...)

    stringslice := strings.Join(contentSlice, " ")
    text, err := sed(stringslice, m.ReferencedMessage.Content)
    if err != nil {
      s.ChannelMessageSendReply(m.ChannelID, "Error: " + err.Error(), m.MessageReference)
      return
    }

    if text == "" {
      text = "[Empty message]"
    }
    s.ChannelMessageSendReply(m.ChannelID, text, m.ReferencedMessage.Reference())
  }
}
