<template>
  <v-row>
    <v-col cols="12" class="text-right">

      <v-btn depressed color="primary" :to="'/books/new'">
        Add Book
        <v-icon>mdi-plus</v-icon>
      </v-btn>
    </v-col>
    <v-col cols="12">
      <v-data-table :headers="headers" :items="dataBuku" :items-per-page="5" class="elevation-1">
        <template v-slot:[`item.actions`]="{ item }">
          <v-btn class="mx-2" fab dark small color="primary" @click="editBook(item)">
            <v-icon>mdi-pencil</v-icon>
          </v-btn>
          <v-btn class="mx-2" fab dark small color="primary" @click="deleteBook(item)">
            <v-icon>mdi-trash-can</v-icon>
          </v-btn>
        </template>
      </v-data-table>
    </v-col>
  </v-row>
</template>

<script>
// import axios from '@nuxtjs/axios'


export default {
  name: 'InspirePage',
  data() {
    return {
      headers: [
        {
          text: 'Title',
          align: 'start',
          sortable: false,
          value: 'Title',
        },
        { text: 'Author', value: 'Author' },
        { text: 'Sypnosis', value: 'Sypnosis' },
        { text: 'Release', value: 'Release' },
        { text: '', value: 'actions' },
      ],
      dataBuku: [],
    }
  },
  mounted() {
    this.$axios.$get('/api/books')
      .then(response => {
        this.dataBuku = response
        console.log(response)
      })
    // this.dataBuku = data.data
  },
  methods: {
    editBook(v) {
      this.$router.push('/books/' + v._id)
    },
    deleteBook(v) { },

  }

}
</script>