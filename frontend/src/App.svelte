<script lang="ts">
  import logo from "./assets/images/logo-universal.png";
  import { CreateEier, GetEiere } from "../wailsjs/go/main/App.js";
  import { models } from "../wailsjs/go/models";

  let resultText = $state("Please enter your name below 👇");
  let name = $state("World");
  let eiere: models.Eier[] = $state([]);

  const fetchEiere = () => {
    GetEiere().then((result) => {
      console.log(result);
      eiere = result;
    });
  };

  const createEier = () => {
    CreateEier(name).then((result) => {
      console.log(result);
      resultText = result.toString();
      fetchEiere();
    });
  };

  fetchEiere();
</script>

<main>
  <div class="container">
    <div class="result" id="result">{resultText}</div>
    {#if eiere.length > 0}
      <div class="result" id="eiere">
        Eier: {eiere.map((eier) => eier.navn).join(", ")}
      </div>
    {/if}
    <div class="input-box" id="input">
      <input
        autocomplete="off"
        bind:value={name}
        class="input"
        id="name"
        type="text"
      />
      <button class="btn" onclick={createEier}>Greet</button>
    </div>
  </div>
</main>

<style>
  main {
    background-color: #282c34;
  }

  .container {
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    color: #fff;
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
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
