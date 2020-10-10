<template>
    <div id="app">
        <!--<HelloWorld msg="Welcome to Your Vue.js App"/>-->

        <!--<component v-bind:is="data.component">-->
        <component v-for="d in data.child" v-bind:key="d.component" v-bind:is="d.component"></component>
        <!--</component>-->

        <draggable v-model="myArray" :clone="dropComponent" @change="log">
            <transition-group>
                <div :key="11">
                    test
                </div>
            </transition-group>
        </draggable>

      <draggable
              class="dragArea list-group"
              :list="list1"
              :group="{ name: 'people', pull: 'clone', put: false }"
              :clone="dropComponent"
              @change="log"
      >
        <div class="list-group-item" v-for="element in list1" :key="element.id">
          {{ element.name }}
        </div>
      </draggable>

        <!--放置区-->
        <div class="drop-container">
            <draggable :list="list2" @change="log">
                <transition-group>
                  <component v-for="d in list2" v-bind:key="d.id" v-bind:is="d.component"></component>
                </transition-group>
            </draggable>
        </div>

    </div>
</template>

<script>
    import HelloWorld from './components/HelloWorld.vue'
    import TextInput from './components/TextInput.vue'
    import draggable from 'vuedraggable'

    // Vue.component('test',{template: '<div>我是一个组件</div>'})

    let data = {
        "component": "HelloWorld",
        child: [
            {
                "component": "HelloWorld"
            },
            {
                "component": "TextInput"
            },
            {
                "component": "TextInput"
            }
        ]
    }

    export default {
        name: 'App',
        components: {
            HelloWorld, TextInput, draggable
        }
        , data() {
            return {
                data: data,
                myArray: null,
                cdata: [],
                list1: [
                    {name: "dog 1", id: 1},
                    {name: "dog 2", id: 2},
                    {name: "dog 3", id: 3},
                    {name: "dog 4", id: 4}
                ],
                list2: [
                    {component: "TextInput", id: 5},
                    {component: "HelloWorld", id: 6},
                    {component: "HelloWorld", id: 7}
                ]
            }

        }
        , methods: {
            dropComponent: function (e) {
                console.log("dropComponent")
                console.log(e)
              this.list2.push({component: "HelloWorld", id: e.id})
            },
            log: function (e) {
                console.log(e)
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

    .drop-container {
        height: 1024px;
        width: 1400px;
        border: 1px solid #2c3e50;
    }
</style>
