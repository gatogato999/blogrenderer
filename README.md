# Blogrenderer

- it takes a post data and generate HTML from it for web server to return in
  response to HTTP requests.
- We'll generate two kinds of page:

1. `View post`: Renders a specific post. The Body field in Post is a string
   containing Markdown so that should be converted to HTML.
2. `Index`: Lists all the posts, with hyperlinks to view the specific post.

## Notes

1. used [ go-approval-test ](https://github.com/approvals/go-approval-tests)
   which allows for easy testing of larger objects, strings and anything else
   that can be saved to a file.
