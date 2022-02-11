package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"strings"
)

// ErrNoAvatar is the error that is returned when the
// Avatar instance is unable to provide an avatar URL.
var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL")
// Avatar represents types capable of representing
// user profile pictures
type Avatar interface {
	// GetAvatarURL gets the avatar URL for the specified client
	// or returns an error if something goes wrong.
	// ErrNoAvatarURL is returned if the object is unable to get
	// a URL for the specified client
	GetAvatarURL(c *client) (string, error)
}

type AuthAvatar struct {}
var UseAuthAvatar AuthAvatar
// func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
// 	if url, ok := c.userData["avatar_url"]; ok {
// 		if urlStr, ok := url.(string); ok {
// 			return urlStr, nil
// 		}
// 	}
// 	return "", ErrNoAvatarURL
// }
func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	url, ok := c.userData["avatar_url"]
	if !ok {
		return "", ErrNoAvatarURL
	}
	urlStr, ok := url.(string)
	if !ok {
		return "", ErrNoAvatarURL
	}
	return urlStr, nil
}


type GravatarAvatar struct {}
var UseGravatar GravatarAvatar
func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	// if email, ok := c.userData["email"]; ok {
	// 	if emailStr, ok := email.(string); ok {
	// 		m := md5.New()
	// 		io.WriteString(m, strings.ToLower(emailStr))
	// 		return fmt.Sprintf("//www.gravatar.com/avatar/%x", m.Sum(nil)), nil
	// 	}
	// }
	// return "", ErrNoAvatarURL
	email, ok := c.userData["email"]
	if !ok {
		return "", ErrNoAvatarURL
	}
	emailStr, ok := email.(string)
	if !ok {
		return "", ErrNoAvatarURL
	}
	m := md5.New()
	io.WriteString(m, strings.ToLower(emailStr))
	return fmt.Sprintf("//www.gravatar.com/avatar/%x", m.Sum(nil)), nil
}