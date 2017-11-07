import React from "react";
import {observer} from 'mobx-react';

import dataStore from './stores/dataStore';

import App from './components/app';
import Header from './components/header';
import Searchbar from './components/search';

@observer
class Main extends React.Component {

  render() {
    var appFilter = new RegExp(dataStore.filter.toLowerCase())

    var displayApps = dataStore.apps.map((app) => {
      if (appFilter.test(app.name.toLowerCase())){
        return app;
      }

      for(var i = 0; i < app.paths.length ; i++ ){
        if (appFilter.test(app.paths[i])){
          return app;
        }
      }

      return undefined;
    }).filter((app) => { return app });

    return (
      <main className="skeleton">
        <article className="container">
          <Header />
          <Searchbar />

          <div className="apps">
            {displayApps.map((app) => {
              return (<App key={app.name} app={app} />)
            })}
          </div>
        </article>
      </main>
    )
  }
}

export default Main
