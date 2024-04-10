https://github.com/sideshow/apns2

# configure the variables in push_server.env

# KeyID from developer account (Certificates, Identifiers & Profiles -> Keys)
APPLE_KEY_ID=KEY_ID
# TeamID from developer account (View Account -> Membership)
APPLE_TEAM_ID=TEAM_ID

APPLE_AUTH_KEY_FILE="cert/AuthKey_xxx.p8"

# build and run in docker
./build.sh && ./run.sh
