<script>
  import { FetchAnyBreed } from "../wailsjs/go/main/App.js";
  import { SelectBreed } from "../wailsjs/go/main/App.js";
  import { FetchByBreed } from "../wailsjs/go/main/App.js";

  let responseObject = {};
  let breeds = [];
  let photos = [];
  let selectedBreed;
  let showRandomPhoto = false;
  let showBreedPhotos = false;

  function init() {
    selectBreed();
  }

  init();

  function fetchAnyBreed() {
    showRandomPhoto = false;
    showBreedPhotos = false;
    FetchAnyBreed().then((result) => (responseObject = result));
    showRandomPhoto = true;
  }

  function selectBreed() {
    SelectBreed().then((result) => (breeds = result));
  }

  function fetchByBreed() {
    init();
    showRandomPhoto = false;
    showBreedPhotos = false;
    FetchByBreed(selectedBreed).then((result) => (photos = result));
    showBreedPhotos = true;
  }
</script>

<main>
  <h3>Dogs API</h3>
  <div>
    <button class="btn" on:click={fetchAnyBreed}
      >Fetch a dog randomly</button
    >
    &nbsp;&nbsp;&nbsp; Click on down arrow to select a breed
    <select bind:value={selectedBreed}>
      {#each breeds as breed}
        <option value={breed}>
          {breed}
        </option>
      {/each}
    </select>
    <button class="btn" on:click={fetchByBreed}>Fetch by this breed</button>
  </div>
  <br />
  {#if showRandomPhoto}
    <img id="random-photo" src={responseObject.message} alt="No dog found" />
  {/if}
  {#if showBreedPhotos}
    {#each photos as photo}
      <img id="breed-photos" src={photo} alt="No dog found" />
    {/each}
  {/if}
</main>

<style>
  #random-photo {
    width: 600px;
    height: auto;
  }

  #breed-photos {
    width: 300px;
    height: auto;
  }

  .btn:focus {
    border-width: 3px;
  }
</style>
