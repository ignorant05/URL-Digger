# URL-Digger 

## Description

A simple DIY URL web-crawler build in golang for learning purposes.
Used as a CLI Tool.

---
 
<details>
<summary><strong>Dependencies</strong></summary>
<br>
- Go v1.24.4
<br>
- [cobra](https://github.com/spf13/cobra).

- [go-pretty](https://github.com/jedib0t/go-pretty?tab=readme-ov-file).
</details>

---

<details>
<summary><strong>Usage</strong></summary>

> Compile it with:

```bash
$ make build 
```

> Then run this to get the description:

```bash
$ ./udig --target <your-target-url> 
```

### Note:
> The output will be a table containing every url found in the anchor html tags in the website .html file.

<br>
</details> 

