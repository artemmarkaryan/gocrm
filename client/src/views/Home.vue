<template>
  <div class="home">
    <h1>Home</h1>
  </div>
</template>

<script>

export default {
  name: 'Home',
  props: {
    serverURL: "http://0.0.0.0:8000"
  }
  data() {
    return {
      users: fetch('https://jsonplaceholder.typicode.com/todos/1')
      .then(response => response.json())
      .then(json => console.log(json))
    }
  }
  components: {
  }
}
</script>
