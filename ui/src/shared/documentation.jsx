import React from "react";
import _ from 'lodash';
import {observer} from 'mobx-react';

import dataStore from './stores/dataStore';

import Swagger from './components/swagger';

@observer
class Documentation extends React.Component {

  render() {
    var name = this.props.match.params.name
    var app = _.find(dataStore.apps, {name: name})

    return (
         <main>
            <Swagger url={app.api} />
         </main>
    )
  }
}

export default Documentation
