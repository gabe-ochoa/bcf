package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Printf("Starting BIG Cat Facts server")

	bigCatFacts := []string{
		"The largest population of tigers in the world is in the USA. There are more captive tigers in the USA alone than there are wild tigers in the entire world.",
		"A male tiger in India adopted a litter of orphaned cubs, taking on the role of ‘mother’. Wildlife officials said such behavior had never been observed before.",
		"Siberian tigers are reported to be the most patient and punishing predators",
		"You shouldn’t use Calvin Klein’s “Obsession” perfume when you are on a Safari. Cheetahs, Tigers, Jaguars and other wild cats are attracted by it. Scientists use it to count the population of these animals better.",
		"Tigers in China are bred in captivity. Those in captivity far outnumber the amount of wild tigers, as tigresses are made to produce cubs at 3 times their natural rate by having cubs weaned off to pigs or dogs. All this is done to make Tiger Bone wine",
		"One leopard in India killed 400 people and was hunted down by the same man who killed a tiger with 436 victims. They were active in the same region at the same time.",
		"The roar from the T-Rex in Jurassic Park is a baby elephant mixed with a tiger and an alligator, and its breath is a whale’s blow.",
		"As an unintended side effect of the Korean Demilitarized Zone, a human free nature preserve has been created that is now acting as a valuable habitat to many endangered species, such as the Korean Tiger and Red Crowned Crane.",
		"In 2005, an Ethiopian girl was kidnapped and beaten by 7 men until a pride of lions chased her attackers off. The lions then stayed and defended her until help arrived.", "A family in Texas in 1953 had a pet lion named Blondie. They never had any incidents with her and she died of old age, as a beloved member of the family.",
		"A “Panther” isn’t a real animal. It is actually a term used to describe black cougars, jaguars and leopards.", "Many lion sanctuaries are actually fronts for canned hunting operations. By letting volunteers touch the animals, they make them easier to eventually kill.",
		"Ostriches sometimes kill lions instead of trying to outrun them.",
		"As late as the early 1900’s jaguar’s natural habitat extended into Texas and almost to Louisiana. There has been various sightings since, including one sighting during 2011 in Arizona.",
		"Greek and Roman word for giraffe was ‘camelopardalis’, because they believed it was an unnatural hybrid of a camel and a leopard.",
		"During their mating season, snow leopards mate 12-36 times a day.",
		"The cougar holds the Guinness record for the animal with the highest number of names, with over 40 in English alone.",
		"Cheetahs were almost wiped out by the last ice-age, and all modern cheetahs are descendants from a small portion of the surviving cats that interbred to maintain their species. Because of this, cheetahs are practically genetic clones of one another.",
		"Dogs are placed with cheetahs in captivity to help keep them calm.",
		"Cheetahs can only run for ~30 seconds before their brain overheats and shuts down",
		"Ancient Egyptians used tamed and trained cheetahs for hunting. This tradition was passed on and kept alive until the 20th century by Indian kings.",
		"A giraffe’s kick is so powerful it can decapitate an African lion.",
		"Cheetah can cover 20+ feet in just one stride!",
		"Snow Leopards can’t purr or roar.",
		"In order to capture footage of tigers hunting in the wild, instead of using a human camera crew, the BBC gave cameras to a team of elephants and trained them to operate the film equipment.",
		"Lions with black manes are alpha males. A dark-maned lion’s testosterone levels are higher than the levels in other males, his cubs are more likely to survive, and he is more likely to recover from wounds. Lionesses prefer a dark-maned mate.",
		"People in the Sundarban Tiger Reserve of Bengal wear masks on the back of their heads to prevent tiger attacks since tigers only attack from behind.", "Male lions lose their mane when neutered",
		"The price of admission for a zoo in 18th century England was a dog or a cat, which were fed to the lions.",
		"Big cats tend to respond to catnip in much the same way as a domestic housecat.",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(time.Now().Unix())
		catFact := bigCatFacts[rand.Intn(len(bigCatFacts))]

		log.Infof("The cat fact is: %s", catFact)
		w.Write([]byte(catFact))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = ":" + port

	http.ListenAndServe(port, nil)
}
