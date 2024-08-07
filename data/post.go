package data

import "time"

type Post struct {
	ID        int
	Title     string
	Content   string
	Timestamp time.Time
}

func GetPosts() []Post {
	return []Post{
		{ID: 0, Title: "What is Lorem IpsumWhat is Lorem IpsumWhat is Lorem IpsumWhat is Lorem Ipsum", Content: ipsum1, Timestamp: time.Now()},
		{ID: 1, Title: "Where does it come from", Content: ipsum2, Timestamp: time.Now()},
		{ID: 2, Title: "Why do we use it", Content: ipsum3, Timestamp: time.Now()},
		{ID: 3, Title: "Where can I get some", Content: ipsum4, Timestamp: time.Now()},
	}
}

const ipsum1 = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam semper a nibh vel sodales. Nam pretium sollicitudin sem et varius. Vestibulum tempor tortor et nunc sollicitudin, id scelerisque velit tincidunt. Suspendisse hendrerit massa id ipsum hendrerit malesuada. Morbi malesuada augue id lacus rutrum pulvinar. Nunc rhoncus sed odio sit amet congue. Proin justo ipsum, elementum eu sodales eu, ullamcorper quis ex. Phasellus finibus neque eu eros semper, ut varius nulla sodales. Praesent tristique, ipsum ac elementum congue, orci elit malesuada risus, in bibendum nunc nibh in libero.\nInteger auctor nunc risus, eu viverra enim cursus vitae. Aenean sed vulputate diam. Praesent vel lacus id nunc commodo finibus. Nulla eu sapien odio. Praesent posuere tincidunt leo, eget vulputate quam. Sed fringilla odio eget finibus auctor. Curabitur scelerisque tempor libero, nec sagittis dui pretium ac. Sed tincidunt sem ac dolor porttitor, ut dictum lorem suscipit. Maecenas eu mollis velit. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Cras interdum, felis at aliquam pharetra, lorem elit dictum felis, quis finibus nisl augue quis felis."
const ipsum2 = "Integer auctor nunc risus, eu viverra enim cursus vitae. Aenean sed vulputate diam. Praesent vel lacus id nunc commodo finibus. Nulla eu sapien odio. Praesent posuere tincidunt leo, eget vulputate quam. Sed fringilla odio eget finibus auctor. Curabitur scelerisque tempor libero, nec sagittis dui pretium ac. Sed tincidunt sem ac dolor porttitor, ut dictum lorem suscipit. Maecenas eu mollis velit. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Cras interdum, felis at aliquam pharetra, lorem elit dictum felis, quis finibus nisl augue quis felis."
const ipsum3 = "Donec vel est scelerisque, pretium justo venenatis, feugiat metus. Phasellus sit amet dolor rutrum, ultrices nulla malesuada, rutrum neque. Curabitur vel orci tortor. Ut egestas elit eros, at luctus eros ullamcorper vel. Suspendisse venenatis vel justo non consequat. Maecenas suscipit condimentum convallis. Phasellus at turpis maximus massa consequat dignissim ac ac ante. Morbi sem eros, ornare sed sapien a, condimentum luctus dolor. Nunc at ornare ante. Phasellus massa lorem, hendrerit ac maximus vel, tempor eget ipsum. Etiam et libero tellus. Vivamus commodo nec est in aliquam. Morbi facilisis magna at euismod ornare. Nullam nec venenatis nibh."
const ipsum4 = "Proin maximus lorem eu ipsum euismod tempor. Nullam ullamcorper erat at nisi blandit, non congue metus suscipit. Fusce consectetur nulla mi, sed molestie nisl mollis quis. Ut erat nunc, convallis in fringilla eu, convallis at metus. Nunc vehicula sem turpis, eget porta mauris elementum porta. Duis maximus nulla sit amet bibendum finibus. Fusce vitae nisl id ante feugiat vehicula. Nunc volutpat eget felis id blandit. Nam efficitur tincidunt egestas. Fusce lacinia, arcu et elementum iaculis, ipsum magna finibus turpis, non convallis neque felis sit amet lectus. Nullam congue laoreet fringilla. Mauris fringilla fringilla lacus, sed varius nisi cursus at. Aliquam id risus hendrerit, rutrum tortor id, ullamcorper nulla. Aenean aliquam urna ut sem porta, in faucibus mauris iaculis."