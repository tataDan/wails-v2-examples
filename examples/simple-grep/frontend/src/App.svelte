<script>
  import { SelectFolder } from "../wailsjs/go/main/App.js";
  import { SelectFile } from "../wailsjs/go/main/App.js";
  import { Search } from "../wailsjs/go/main/App.js";

  let path;
  let pattern;
  let caseInsensitive = false;
  let wholeWord = false;
  let wholeLine = false;
  let filenameOnly = false;
  let filesWoMatches = false;
  let options = {
    caseInsensitive: caseInsensitive,
    wholeWord: wholeWord,
    wholeLine: wholeLine,
    filenameOnly: filenameOnly,
    filesWoMatches: filesWoMatches,
  };
  let results = "";

  function processOptions(event) {
    options = {
      caseInsensitive: caseInsensitive,
      wholeWord: wholeWord,
      wholeLine: wholeLine,
      filenameOnly: filenameOnly,
      filesWoMatches: filesWoMatches,
    };
  }

  function selectFolder() {
    SelectFolder().then((result) => (path = result));
  }

  function selectFile() {
    SelectFile().then((result) => (path = result));
  }

  function search() {
    results = "Searching...";
    Search(path, pattern, options).then((result) => (results = result));
  }
</script>

<main>
  <h3 class="title">SIMPLE GREP</h3>
  <div class="input-box" id="input">
    <label for="path">Path:</label>
    &nbsp;
    <input
      autocomplete="off"
      bind:value={path}
      class="input"
      id="path"
      type="text"
    />
    &nbsp;&nbsp;&nbsp;
    <label for="pattern">Pattern:</label>
    &nbsp;
    <input
      autocomplete="off"
      bind:value={pattern}
      class="input"
      id="pattern"
      type="text"
    />
    <br /><br />
  </div>
  <div class="checkboxes">
    <span class="checkboxes-with-labels">
      <input
        class="checkbox"
        type="checkbox"
        id="caseInsensitive"
        value="caseInsensitive"
        bind:checked={caseInsensitive}
        on:change={processOptions}
      />
      <label for="caseInsensitive">Case Insensitive</label>
    </span>
    <span class="checkboxes-with-labels">
      <input
        class="checkbox"
        type="checkbox"
        id="wholeWord"
        value="wholeWord"
        bind:checked={wholeWord}
        on:change={processOptions}
      />
      <label for="wholeWord">Whole Word</label>
    </span>
    <span class="checkboxes-with-labels">
      <input
        class="checkbox"
        type="checkbox"
        id="wholeLine"
        value="wholeLine"
        bind:checked={wholeLine}
        on:change={processOptions}
      />
      <label for="wholeLine">Whole Line</label>
    </span>
    <span class="checkboxes-with-labels">
      <input
        class="checkbox"
        type="checkbox"
        id="fileNameOnly"
        value="fileNameOnly"
        bind:checked={filenameOnly}
        on:change={processOptions}
      />
      <label for="fileNameOnly">File Name Only</label>
    </span>
    <span class="checkboxes-with-labels">
      <input
        class="checkbox"
        type="checkbox"
        id="filesWoMatches"
        value="filesWoMatches"
        bind:checked={filesWoMatches}
        on:change={processOptions}
      />
      <label for="filesWoMatches">Files Without Matches</label>
    </span>
  </div>
  <br />
  <div class="buttons">
    <button class="btn" id="select-folder-btn" on:click={selectFolder}
      >Select a Folder</button
    >
    <button class="btn" id="select-file-btn" on:click={selectFile}
      >Select a File</button
    >
    <button class="btn" id="search-btn" on:click={search}>Search</button>
  </div>
  <br />
  <div>
    <textarea value={results} />
  </div>
</main>

<style>
  .title {
    margin: 10px;
  }

  .buttons {
    display: flex;
    justify-content: flex-start;
  }

  .btn {
    margin-left: 10px;
    margin-right: 10px;
  }

  .btn:focus {
    border-width: 3px;
  }

  .checkbox:focus {
    outline: 2px solid white;
  }

  .checkboxes {
    display: flex;
    justify-content: flex-start;
  }

  .checkboxes-with-labels {
    margin-left: 10px;
    margin-right: 10px;
  }

  textarea {
    width: 100%;
    height: 550px;
    background-color: gray;
  }

  #pattern {
    width: 250px;
  }

  .input-box {
    display: flex;
    justify-content: flex-start;
    padding: 0px 10px;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    width: 450px;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }
</style>
