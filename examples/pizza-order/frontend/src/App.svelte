<script>
  import { Order } from "../wailsjs/go/main/App";

  let resultText = "";
  let name;
  let phone;
  let address;

  let size = "small";
  let small = true;
  let medium = false;
  let large = false;

  let crust = "thin";
  let thin = true;
  let thick = false;

  let pepperoni = false;
  let sausage = false;
  let linguica = false;
  let anchovies = false;
  let olives = false;
  let mushrooms = false;
  let tomatoes = false;

  let options = {
    small: small,
    medium: medium,
    large: large,
    thin: thin,
    thick: thick,
    pepperoni: pepperoni,
    sausage: sausage,
    linguica: linguica,
    anchovies: anchovies,
    olives: olives,
    mushrooms: mushrooms,
    tomatoes: tomatoes,
  };

  function processRadioAndCheckValues(event) {
    console.log("Hello");
    console.log(size);
    if (size === "small") {
      small = true;
      medium = false;
      large = false;
    } else if (size === "medium") {
      small = false;
      medium = true;
      large = false;
    } else if (size === "large") {
      small = false;
      medium = false;
      large = true;
    }
    console.log(crust);
    if (crust === "thin") {
      thin = true;
      thick = false;
    } else if (crust === "thick") {
      thin = false;
      thick = true;
    }

    console.log(small);
    console.log(medium);
    console.log(large);
    console.log(thin);
    console.log(thick);

    options = {
      small: small,
      medium: medium,
      large: large,
      thin: thin,
      thick: thick,
      pepperoni: pepperoni,
      sausage: sausage,
      linguica: linguica,
      anchovies: anchovies,
      olives: olives,
      mushrooms: mushrooms,
      tomatoes: tomatoes,
    };
  }

  function order() {
    Order(name, phone, address, options).then(
      (result) => (resultText = result)
    );
  }
</script>

<main>
  <h3>ORDER YOUR PIZZA NOW</h3>
  <div class="input-box" id="input">
    <label for="name">Name:</label>
    <input
      autocomplete="off"
      bind:value={name}
      class="input"
      id="name"
      type="text"
    />
    <br /><br />
    <label for="phone">Phone:</label>
    <input
      autocomplete="off"
      bind:value={phone}
      class="input"
      id="phone"
      type="text"
    />
    <br /><br />
    <label for="address">Address:</label>
    <input
      autocomplete="off"
      bind:value={address}
      class="input"
      id="address"
      type="text"
    />
    <br /><br />

    <div class="size">
      <u>Size</u>
      <br />
      <input
        type="radio"
        id="small"
        class="radio"
        value="small"
        name="size"
        bind:group={size}
        on:change={processRadioAndCheckValues}
      />
      <label for="small">Small</label><br />
      <input
        type="radio"
        id="medium"
        class="radio"
        value="medium"
        name="size"
        bind:group={size}
        on:change={processRadioAndCheckValues}
      />
      <label for="medium">Medium</label><br />
      <input
        type="radio"
        id="large"
        class="radio"
        value="large"
        name="size"
        bind:group={size}
        on:change={processRadioAndCheckValues}
      />
      <label for="large">Large</label>
    </div>

    <div class="crust">
      <u>Crust</u>
      <br />
      <input
        type="radio"
        class="radio"
        id="thin"
        value="thin"
        name="crust"
        bind:group={crust}
        on:change={processRadioAndCheckValues}
      />
      <label for="thin">Thin</label>
      <br />
      <input
        type="radio"
        id="thick"
        class="radio"
        value="thick"
        name="crust"
        bind:group={crust}
        on:change={processRadioAndCheckValues}
      />
      <label for="thick">Thick</label>
    </div>

    <div class="toppings">
      <u>Toppings</u>
      <br /><br />
      <span class="checkboxes-with-labels">
        <input
          type="checkbox"
          id="pepperoni"
          class="checkbox"
          value="Pepperoni"
          bind:checked={pepperoni}
          on:change={processRadioAndCheckValues}
        />
        <label for="pepperoni">Pepperoni</label>
      </span>
      <span class="checkboxes-with-labels">
        <input
          type="checkbox"
          id="sausage"
          class="checkbox"
          value="Sausage"
          bind:checked={sausage}
          on:change={processRadioAndCheckValues}
        />
        <label for="sausage">Sausage</label>
        <span class="checkboxes-with-labels">
          <input
            type="checkbox"
            id="linguica"
            class="checkbox"
            value="Linguica"
            bind:checked={linguica}
            on:change={processRadioAndCheckValues}
          />
          <label for="linguica">Linguica</label>
        </span>
        <span class="checkboxes-with-labels">
          <input
            type="checkbox"
            id="anchovies"
            class="checkbox"
            value="Anchovies"
            bind:checked={anchovies}
            on:change={processRadioAndCheckValues}
          />
          <label for="anchovies">Anchovies</label>
          <br /><br />
        </span>
        <span class="checkboxes-with-labels">
          <input
            type="checkbox"
            id="olives"
            class="checkbox"
            value="Olives"
            bind:checked={olives}
            on:change={processRadioAndCheckValues}
          />
          <label for="olives">Olives</label>
        </span>
        <span class="checkboxes-with-labels">
          <input
            type="checkbox"
            id="mushrooms"
            class="checkbox"
            value="Mushrooms"
            bind:checked={mushrooms}
            on:change={processRadioAndCheckValues}
          />
          <label for="mushrooms">Mushrooms</label>
        </span>
        <span class="checkboxes-with-labels">
          <input
            type="checkbox"
            id="tomatoes"
            class="checkbox"
            value="Tomatoes"
            bind:checked={tomatoes}
            on:change={processRadioAndCheckValues}
          />
          <label for="tomatoes">Tomatoes</label>
        </span>
      </span>
    </div>

    <br /><br /><br /><br /><br /><br /><br />

    <button class="btn" on:click={order}>Submit</button>
    <br /><br />
  </div>

  <textarea value={resultText} />
</main>

<style>
  .checkboxes-with-labels {
    margin-left: 10px;
    margin-right: 10px;
  }

  .checkbox:focus {
    outline: 2px solid white;
  }

  .crust {
    position: absolute;
    left: 275px;
    width: 100px;
    border: 2px solid black;
    padding: 10px;
  }

  .radio {
    margin-top: 10px;
    margin-bottom: 8px;
  }

  .radio:focus {
    outline: 1px solid white;
  }

  textarea {
    width: 95%;
    height: 150px;
    background-color: gray;
  }

  .input-box .btn {
    width: 80px;
    height: 30px;
    line-height: 30px;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:focus {
    border-width: 3px;
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

  .size {
    position: absolute;
    left: 100px;
    width: 100px;
    border: 2px solid black;
    padding: 10px;
  }

  .toppings {
    position: absolute;
    left: 450px;
    width: 450px;
    border: 2px solid black;
    padding: 10px;
  }
</style>
