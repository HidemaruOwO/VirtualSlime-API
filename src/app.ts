import express, { Request, Response, NextFunction } from "express";
import fs from "fs";
import path from "path";
import matter from "gray-matter";

// Types imports
import { Frontmatter, Post } from "./interface";

require("dotenv").config();

const app = express();

app.use(express.json());

console.log("🔍 Checking $NODE_ENV");
if (process.env.NODE_ENV === "production") {
  console.log("😀 This is Production");
} else {
  console.log("🛠️ This is Development");
}

app.use((req: Request, res: Response, next: NextFunction) => {
  if (process.env.NODE_ENV === "production") {
    res.header("Access-Control-Allow-Origin", process.env.CORS_ALLOW_ORIGIN);
    res.header("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE");
    res.header("Access-Control-Allow-Headers", "*");
  } else {
    res.header("Access-Control-Allow-Origin", "*");
    res.header("Access-Control-Allow-Methods", "*");
    res.header("Access-Control-Allow-Headers", "*");
  }
  next();
});

app.get("/", (req: Request, res: Response) => {
  res.status(200).send("Running Virtual Slime API");
});

app.get("/api/posts", (req: Request, res: Response) => {
  let posts: Post[] = [];

  if (process.env.NODE_ENV === "production") {
    // Fetch from cache
    posts = require("../cache/data").posts;
  } else {
    const postDir = path.join(__dirname, "..", "posts", "posts");

    console.log(postDir);

    const files = fs.readdirSync(postDir);

    posts = files.map((filename) => {
      const slug = filename.replace(".md", "");

      const markdownWithMeta = fs.readFileSync(
        path.join(postDir, filename),
        "utf-8"
      );

      const { data: frontmatter } = matter(markdownWithMeta);

      return {
        slug,
        frontmatter: frontmatter as Frontmatter, // Explicitly cast to Frontmatter type
      };
    });
  }

  const query = (req.query.q as string).toLowerCase();

  const results = posts.filter(
    ({ frontmatter: { title, excerpt, category } }) =>
      title.toLowerCase().includes(query) ||
      excerpt.toLowerCase().includes(query) ||
      category.toLowerCase().includes(query)
  );

  res.status(200).json({ results });
});

const PORT: number = parseInt(process.env.PORT as string, 10) || 3000;

app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}`);
});
