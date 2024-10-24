// this is a simple markdown cheatsheet, which can be used to help formatting
// note that this syntax displayed in the sheet only covers the basics

package utils

const Cheatsheet = `
MARKDOWN SYNTAX CHEATSHEET

1. HEADERS
   # H1
   ## H2
   ### H3
   #### H4
   ##### H5
   ###### H6

   Alternatively, for H1 and H2:
   
   Header 1
   ========
   
   Header 2
   --------

2. EMPHASIS
   *italic*  or _italic_
   **bold**  or __bold__
   **_bold italic_**
   ~~strikethrough~~

3. LISTS
   Unordered:
   * Item 1
   * Item 2
     * Subitem 2.1
     * Subitem 2.2

   Ordered:
   1. First item
   2. Second item
   3. Third item

4. LINKS
   [Link text](URL "Optional title")
   [Reference-style link][reference]
   [reference]: URL "Optional title"

5. IMAGES
   ![Alt text](URL "Optional title")
   ![Alt text][image-reference]
   [image-reference]: URL "Optional title"

6. BLOCKQUOTES
   > This is a blockquote
   > It can span multiple lines

7. CODE
   Inline code: ` + "`code`" + `
   
   Code block:
       ` + "```" + `language
       code block
       ` + "```" + `
   Or indent each line by 4 spaces

8. HORIZONTAL RULE
   Three or more hyphens, asterisks, or underscores:
   ---
   ***
   ___

9. TABLES
   | Header 1 | Header 2 |
   |----------|----------|
   | Cell 1   | Cell 2   |
   | Cell 3   | Cell 4   |

10. TASK LISTS
    - [x] Completed task
    - [ ] Incomplete task

11. ESCAPING CHARACTERS
    Use a backslash to escape special characters:
    \*This will show asterisks\*
`
