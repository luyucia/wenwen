<template>
  <div id="app">
    <!--<HelloWorld msg="Welcome to Your Vue.js App"/>-->

    <!--<component v-bind:is="data.component">-->
      <component v-for="d in data.child" v-bind:key="d.component" v-bind:is="d.component"></component>
    <!--</component>-->

    <draggable v-model="myArray">
      <transition-group>
        <div :key="11">
          test
        </div>
      </transition-group>
    </draggable>

  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import TextInput from './components/TextInput.vue'
import draggable from 'vuedraggable'

// Vue.component('test',{template: '<div>我是一个组件</div>'})

let data = {
    "component":"HelloWorld",
    child:[
        {
            "component":"HelloWorld"
        },
        {
            "component":"TextInput"
        },
        {
            "component":"TextInput"
        }
    ]
}

export default {
  name: 'App',
  components: {
    HelloWorld,TextInput,draggable
  },data(){
        return{
            data:data,
            myArray:null
        }
  }

}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
