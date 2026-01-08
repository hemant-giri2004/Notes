<h1 style="border-bottom:2px solid #e5e7eb;padding-bottom:8px;">
📝 Notes Application (Full-Stack)
</h1>

<p>
A full-stack <strong>Notes Application</strong> built using
<strong>Angular</strong> for the frontend,
<strong>Go (Golang)</strong> with <strong>Gorilla Mux</strong> for the backend,
and <strong>PostgreSQL</strong> for data storage.
</p>

<p>
This project demonstrates real-world full-stack development including
REST APIs, database integration, environment configuration, CORS handling,
and cloud deployment.
</p>

<hr />

<h2>🚀 Live Demo</h2>

<ul>
  <li>
    <strong>Frontend:</strong>
    <a href="https://hemant-giri2004.github.io/Notes/" target="_blank">
      GitHub Pages
    </a>
  </li>
  <li>
    <strong>Backend API:</strong>
    <a href="https://notes-backend-t4bj.onrender.com" target="_blank">
      Render
    </a>
  </li>
</ul>

<hr />

<h2>🛠 Tech Stack</h2>

<ul>
  <li>🌐 Angular (Standalone Architecture)</li>
  <li>⚙️ Go (Golang)</li>
  <li>🧭 Gorilla Mux</li>
  <li>🗄 PostgreSQL (Neon)</li>
  <li>☁️ GitHub Pages (Frontend)</li>
  <li>☁️ Render (Backend)</li>
</ul>

<hr />

<h2>✨ Features</h2>

<ul>
  <li>➕ Add a new note</li>
  <li>📄 View all notes</li>
  <li>❌ Delete a note</li>
  <li>🔄 Frontend ↔ Backend communication</li>
  <li>🌍 Environment-based configuration</li>
</ul>

<hr />

<h2>📂 Project Structure</h2>

<pre style="
background:#0f172a;
color:#e5e7eb;
padding:16px;
border-radius:6px;
overflow-x:auto;
">
Notes/
├── backend/
│   ├── main.go        # App entry point
│   ├── cors.go        # CORS configuration
│   ├── db/            # Database connection
│   ├── handlers/      # API handlers
│   └── models/        # Data models
│
├── frontend/
│   ├── src/
│   │   ├── app/
│   │   │   ├── pages/     # UI components
│   │   │   ├── services/  # API services
│   │   │   └── app.ts
│   │   ├── environments/
│   │   └── main.ts
│   ├── angular.json
│   └── package.json
│
├── docs/              # Production build (GitHub Pages)
│
└── README.md
</pre>

<hr />

<h2>🔗 API Endpoints</h2>

<p><strong>Base URL:</strong></p>

<pre>
https://notes-backend-t4bj.onrender.com
</pre>

<ul>
  <li><code>GET /notes</code> – Fetch all notes</li>
  <li><code>POST /notes</code> – Add a new note</li>
  <li><code>DELETE /notes/{id}</code> – Delete a note</li>
</ul>

<hr />

<h2>🧪 Local Development</h2>

<h3>Backend</h3>
<ul>
  <li>Go 1.21+</li>
  <li>Create <code>.env</code> with <code>DATABASE_URL</code></li>
  <li>Run using <code>go run main.go</code></li>
</ul>

<h3>Frontend</h3>
<ul>
  <li>Node.js + Angular CLI</li>
  <li><code>npm install</code></li>
  <li><code>ng serve</code></li>
</ul>

<hr />

<h2>📦 Deployment</h2>

<ul>
  <li>🌐 Frontend deployed on GitHub Pages</li>
  <li>⚙️ Backend deployed on Render</li>
  <li>🗄 Database hosted on Neon (PostgreSQL)</li>
</ul>

<hr />

<h2>🧠 Key Learnings</h2>

<ul>
  <li>Angular standalone architecture</li>
  <li>REST API development using Go</li>
  <li>Environment handling (local vs production)</li>
  <li>CORS configuration</li>
  <li>Debugging real deployment issues</li>
</ul>

<hr />

<h2>🔮 Future Enhancements</h2>

<ul>
  <li>🔐 Authentication (JWT)</li>
  <li>✏️ Update notes feature</li>
  <li>🎨 Improved UI/UX</li>
  <li>🐳 Docker & CI/CD</li>
</ul>

<hr />

<p>
<strong>👨‍💻 Author:</strong><br />
Hemant Giri<br />
MCA Student | Aspiring Software Engineer
</p>

<p>
📌 This project is built for learning and demonstration purposes.
</p>
