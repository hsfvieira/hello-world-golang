curl -X POST -d "{
  \"name\": \"Henrique Fioreli\",
  \"username\": \"hsfvieira\",
  \"email\": \"henrique@fioreli.dev\"
}" http://localhost:8080/api/users
echo ""
curl -X POST -d "{
  \"name\": \"Henrique Fioreli 2\",
  \"username\": \"hsfvieira2\",
  \"email\": \"henrique2@fioreli.dev\"
}" http://localhost:8080/api/users
echo ""
curl http://localhost:8080/api/users | jq
echo ""
curl http://localhost:8080/api/users/hsfvieira | jq
echo ""
