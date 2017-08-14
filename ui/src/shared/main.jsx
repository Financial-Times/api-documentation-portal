import React from "react";
import {observer} from 'mobx-react';

import dataStore from './stores/dataStore';

import App from './components/app';
import Banner from './components/banner';
import Header from './components/header';
import Sidebar from './components/sidebar';

@observer
class Main extends React.Component {

  render() {
    return (
         <main>
            <Header />
            <Banner />

            <div className="o-techdocs-container">
               <div className="o-techdocs-layout">

                  <Sidebar apps={dataStore.apps} />

                  <div className="o-techdocs-main">
                     <div className="o-techdocs-content">
                        {dataStore.apps.map((app) => {
                          return (<App key={app.name} app={app} />)
                        })}
                     </div>
                  </div>
               </div>
            </div>
         </main>
    )
  }
}

export default Main
