# Go: Channel

## Critical operation on a channel

| Operation | Channel State | Result |
|--|--|--|
|Read|nil|Block|
|Write|nil|Block|
|Write|closed|panic|
|Close|nil|panic|
|Close|closed|panic|

根據上表，當讀取或寫入一個 nil 的 channel 時，該 goroutine 會 block 住，
而剩下更危險的是另外三種情況是會 panic 的

為了解決上述問題，一個原則是建立 channel 的 goroutine 需要也只能是唯一的一個地方可以寫入及關閉這個 channel
範例程式請見 channel.go




