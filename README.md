# Making an HTTP Web Server in Go (from Scratch)

Pretty fun :)

I'm building this server from the ground up in Go as a way to deepen my backend development skills and get more comfortable with industry-relevant fundamentals. It's a work-in-progress, and this README is here both for anyone stopping by and for myself when I return after a break because of school.

---

### 📌 Currently Implemented Routes

##### Sidenote: curently the database is completely in memory of the server (plan to migrate to PostgreSQL once all functionality of server is complete) there are some starter data hard-coded into memory to play around with though

#### **Playlist Routes**

- `GET    {baseURL}/playlist/{name}`  
  → Fetch information about a single playlist
- `GET    {baseURL}/playlist`  
  → Fetch information about all playlists in the database
- `POST   {baseURL}/playlist`  
  → Add a new playlist to the database
- `POST   {baseURL}/playlist/{name}/tracks`  
  → Add a new track to a specific playlist

#### **Track Routes**

- `GET    {baseURL}/tracks/{id}`  
  → Fetch information about a specific track
- `POST   {baseURL}/tracks/{id}`  
  → Upvote a specific track

---

More updates to come as I keep learning and building. Feel free to copy repo and play around with Postman, etc!
