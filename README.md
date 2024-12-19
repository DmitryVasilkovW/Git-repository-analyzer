# **Gitanalyzer**

gitanalyzer is a console utility for counting author statistics in a Git repository. You can use it to analyze who has contributed the most to a project based on lines of code, commits, and files.

## **Example of use**

````bash
./gitanalyzer --repository=git@github.com:DmitryVasilkovW/Git-repository-analyzer.git --revision=HEAD --order-by=commits --show-languages=true
````

**Result:**

````  
User                 Files      Lines      Commits    Languages
Dmitry Vasilkov      88         7335       2          JSON (3377), Go Module (5), Go Checksums (2), YAML (150), Ignore List (117), Makefile (2), Markdown (252), Go (1183)
````

## **Capabilities**

gitanalyzer provides the following metrics for each author:

* **Number of lines** of code
* **Number of unique commits**
* **Number of files** affected by author commits

All statistics are calculated based on the last state of the repository.



## **Flags**

gitanalyzer supports the following flags:

* **`--repository`**  
  Path to the repository (both local and remote, current directory by default).
* **`-revision`**  
  Pointer to commit (default is `HEAD`).
* **`--order-by`**  
  Results sorting key: `lines` (default), `commits`, `files`.
* **`--use-committer`**  
  Replaces author with committer.
* **`--format`**  
  Output format:
  * `tabular` (default).
  * `csv`
  * `json`
  * `json-lines`
* **`--extensions`**  
  List of file extensions to analyze, e.g. ``.go,.md'``.
* **`-languages`**  
  List of languages to analyze, e.g., .go,markdown`.
* **`--exclude`**  
  Glob-patterns to exclude files, e.g., ``foo/*,bar/*'``.
* **`--restrict-to`**  
  Glob-patterns to filter only the files you want.
* **`-show-languages`**  
  if true, the utility will display all used languages for each user and the number of lines for each language

## **Testing**

To run tests, use the command:

````bash
go test -v test/integration/gitanalyzer_test.go
````

Integration tests use pre-built Git bundles stored in `/tests/integration/testdata/bundles`. To create your own bundle, run:

````bash
git bundle create my.bundle --all
````

Unzip the bundle:

````bash
`git clone /path/to/my.bundle .`
````