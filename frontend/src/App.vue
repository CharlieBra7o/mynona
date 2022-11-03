<script>
import axios from "axios";

export default {
  name: "App",

  data() {
    return {
      domains: null,
      domain: null,
      mailboxes: null,
      identity: null,
      aliases: null,
    };
  },

  methods: {
    getDomains() {
      console.log(`CALL getDomains()`);

      // Call the Go API, in this case we only need the URL parameter.
      axios
        .get("http://localhost:3000/api/domains", {})
        .then((response) => {
          this.domains = response.data;
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error}`);
        });
    },

    getMailboxes(domain) {
      console.log("CALL getMailboxes(" + domain + ")");

      // Call the Go API, in this case we only need the URL parameter.
      axios
        .get("http://localhost:3000/api/domains/" + domain + "/mailboxes", {})
        .then((response) => {
          this.mailboxes = response.data.mailboxes;
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error}`);
        });
    },

    getAliases(domain) {
      // Call the Go API, in this case we only need the URL parameter.
      console.log("CALL getAliases(" + domain + ")");

      axios
        .get("http://localhost:3000/api/domains/" + domain + "/aliases", {})
        .then((response) => {
          this.aliases = response.data.address_aliases;
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error}`);
        });
    },

    addIdentity(domain) {
      // Call the Go API, in this case we only need the URL parameter.
      console.log("CALL addIdentity(" + domain + ")");

      axios
        .post(
          "http://localhost:3000/api/domains/" +
            domain +
            "/mailboxes/demo/identities",
          {
            name: "Identity Name",
            local_part: this.identity,
            password: "Sup3r_s3cr3T",
          }
        )
        .then((response) => {
          console.log("[DEBUG] Response: " + response.data);
          this.aliases = response.data;
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error}`);
        });
    },

    onDomainSelected(event) {
      console.log("EVENT onDomainSelected: " + event.target.value);
      this.getMailboxes(event.target.value);
      this.getAliases(event.target.value);
    },
  },

  beforeMount() {
    console.log(`CALL beforeMount()`);
    this.getDomains();
  },
};
</script>

<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>mynona.</h1>

        <div>
          <h2>Domain</h2>
          <select
            name="domain"
            id="domain"
            v-model="domain"
            @change="onDomainSelected($event)"
          >
            <option v-for="domain in domains" :key="domain">
              {{ domain }}
            </option>
          </select>

          <h3>Mailboxes</h3>
          <ul id="mailboxes">
            <li v-for="mailbox in mailboxes" :key="mailbox.address">
              {{ mailbox.local_part }}
              <ul>
                <li
                  v-for="identity in mailbox.identities"
                  :key="identity.address"
                >
                  {{ identity.local_part }} ({{ identity.name }})
                </li>
              </ul>
            </li>
          </ul>

          <h3>Aliases</h3>
          <ul id="aliases">
            <li v-for="alias in aliases" :key="alias.address">
              {{ alias.local_part }}
            </li>
          </ul>
        </div>

        <form v-on:submit.prevent="addIdentity">
          <div class="form-group">
            <input
              v-model="identity"
              type="text"
              id="website-input"
              placeholder="Enter an identity"
              class="form-control"
            />
          </div>
          <div class="form-group">
            <button class="btn btn-primary">Add</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>
