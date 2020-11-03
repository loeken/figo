# Run
DB_NAME=Football DB_USER=root DB_PASS=topsecure DB_HOST=127.0.0.1 DB_MIGRATE=true go run ./main.go

# create a release record
curl -H "Content-Type: application/json" localhost:3000/api/v1/release -X POST --data '{"title": "release #1", "artist": "rZumA"}'  

# get all releases
curl localhost:3000/api/v1/release

# get a release by id
curl localhost:3000/api/v1/release


# delete release
curl localhost:3000/api/v1/release/1 -X DELETE

# update release
curl -H "Content-Type: application/json" localhost:3000/api/v1/release/1 -X PUT  --data '{"title": "real release #1", "artist": "RzumA"}' 

# upload attachement
curl -F 'attachment=@/home/loeken/Pictures/2020-11-03-111529_564x180_scrot.png' localhost:3000/api/v1/release/1/upload