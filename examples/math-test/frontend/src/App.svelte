<script>
  import { TestOne } from "../wailsjs/go/main/App.js";
  import { TestTwo } from "../wailsjs/go/main/App.js";
  import { SubmitAnswer } from "../wailsjs/go/main/App.js";

  import gif1 from "./assets/images/well-done-good-job.gif";
  import gif2 from "./assets/images/giphy.gif";

  let test = [];
  let answer = "";
  let feedback = [];
  let showSubmitBtn = false;
  let perfectScore = false;

  function getRndInteger(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
  }
  let randomNum;

  function testOne() {
    randomNum = getRndInteger(1, 2);
    perfectScore = false;
    feedback = [];
    TestOne().then((result) => (test = result));
    showSubmitBtn = true;
  }

  function testTwo() {
    randomNum = getRndInteger(1, 2);
    perfectScore = false;
    feedback = [];
    TestTwo().then((result) => (test = result));
    showSubmitBtn = true;
  }

  function submitAnswer() {
    SubmitAnswer(test).then((result) => (feedback = result));
  }

  // @ts-ignore
  window.runtime.EventsOn("perfect-score", function () {
    perfectScore = true;
  });
</script>

<main>
  <br />

  <h3>MATH TEST</h3>
  {#if perfectScore}
    {#if randomNum === 1}
      <!-- svelte-ignore a11y-img-redundant-alt -->
      <img alt="Missing Photo" id="logo" src={gif1} />
    {:else}
      <!-- svelte-ignore a11y-img-redundant-alt -->
      <img alt="Missing Photo" id="logo2" src={gif2} />
    {/if}
  {/if}
  <div class="input-box" id="input">
    <button class="btn" on:click={testOne}>Test One</button>
    <button class="btn" on:click={testTwo}>Test Two</button>
    <br /><br />

    {#each test as data, i}
      {data.Id}.&nbsp;
      <label for="question">Question:</label>
      <input
        autocomplete="off"
        bind:value={data.Question}
        class="input"
        id="question"
        type="text"
      />
      <label for="response">Answer:</label>
      <input
        autocomplete="off"
        bind:value={data.Response}
        class="input"
        id="response"
        type="text"
      />

      {#if feedback.length > 0}
        {#if feedback[i]}
          Correct
        {/if}
      {/if}
      <br /><br />
    {/each}
    {#if showSubmitBtn}
      <button class="btn" on:click={submitAnswer}>Submit</button>
    {/if}
  </div>
</main>

<style>
  .btn:focus {
    border-width: 3px;
  }

  .input-box .btn {
    width: 80px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
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
