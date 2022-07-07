<script>
  import { GetCountryNames } from "../wailsjs/go/main/App.js";
  import { GetCountry } from "../wailsjs/go/main/App.js";

  let countryNames = [];
  let selectedCountry;
  let country = {};
  let showCountryData = false;

  function init() {
    getCountryNames();
  }

  init();

  function getCountryNames() {
    GetCountryNames().then((result) => (countryNames = result));
  }

  function getCountry() {
    showCountryData = false;
    GetCountry(selectedCountry).then((result) => (country = result));
    showCountryData = true;
  }
</script>

<h3>Country API</h3>
<div>
  &nbsp;&nbsp;&nbsp; Click on down arrow to select a country
  <select bind:value={selectedCountry}>
    {#each countryNames as cn}
      <option value={cn}>
        {cn}
      </option>
    {/each}
  </select>
  <button class="btn" on:click={getCountry}>Get Country</button>
</div>
{#if showCountryData}
  <h3>{selectedCountry}</h3>
  <ul>
    <li class="flag">Flag: {country.Flag}</li>
    <li>Capital: {country.Capital}</li>
    <li>Cca2: {country.Cca2}</li>
    <li>Cca3: {country.Cca3}</li>
    <li>Region: {country.Region}</li>
    <li>Continents: {country.Continents}</li>
    <li>Languages: {country.LanguageList}</li>
    <li>Population: {country.PopulationText}</li>
    <li>Area: {country.AreaText}</li>
    <li>Currencies: {country.CurrencyList}</li>
    {#if country.UnMember}
      <li>U.N. Member: Yes</li>
    {:else}
      <li>U.N. Member: No</li>
    {/if}
    {#if country.Independent}
      <li>Independent: Yes</li>
    {:else}
      <li>Independent: No</li>
    {/if}
    {#if country.Landlocked}
      <li>Landlocked: Yes</li>
    {:else}
      <li>Landlocked: No</li>
    {/if}
    <li>{country.LatlngText}</li>
    <li>Timezones: {country.Timezones}</li>
    <li>Flags: {country.FlagList}</li>
  </ul>
{/if}

<style>
  ul {
    list-style-type: none;
  }
</style>
