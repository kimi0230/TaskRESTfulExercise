print(
  "Start #################################################################"
);

// db.auth("root", "root");

db = db.getSiblingDB("task-restful-exercise");
db.createUser({
  user: "user",
  pwd: "user_password",
  roles: [{ role: "readWrite", db: "task-restful-exercise" }],
});
db.createCollection("users");

db = db.getSiblingDB("task-restful-exercise");

db.createUser({
  user: "api_user",
  pwd: "api1234",
  roles: [{ role: "readWrite", db: "task-restful-exercise" }],
});
db.createCollection("users");

print("END #################################################################");
