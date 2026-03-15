module k6-game-test

go 1.26.1

require (
	golang.org/x/time v0.15.0
	google.golang.org/protobuf v1.36.11
)

replace k6-game-test/proto/pb => ./proto/pb
