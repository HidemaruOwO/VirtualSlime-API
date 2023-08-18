export type Frontmatter = {
  title: string;
  excerpt: string;
  category: string;
}

export type Post = {
  slug: string;
  frontmatter: Frontmatter;
}


