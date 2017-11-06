import React, { Component } from "react";
import dataStore from '../stores/dataStore';

export default class Search extends Component {
  constructor (props) {
    super(props)
  }

  search = () => {
    dataStore.filter = this.refs.search.value
  }

  render() {
    return (
      <section className="search">
        <label htmlFor="search">Search</label>
        <input id="search"
          ref='search'
          className="u-full-width"
          onKeyUp={this.search}
          type="search"
          placeholder="Application to search for"
        />
      </section>
    );
  }
}
