import React from "react";
import {observer} from "mobx-react";
import {Link} from 'react-router-dom';

/**
 * Component showing progress bar for a single collection cycle
 * @type {Observer}
 */
@observer
class App extends React.Component {

  render(){
    let app = this.props.app;

    return (
         <div>
            <h2 id={app.name}>{app.title}</h2>
            <p>{app.description}</p>
            <div className="o-techdocs-card">
               <div className="o-techdocs-card__context">
                  <div className="o-techdocs-card__heading">
                     <div className="o-techdocs-card__title">API Documentation</div>
                  </div>
                  <div className="o-techdocs-card__quickactions">
                     <Link to={"/docs/" + app.name}><button className="o-techdocs-card__actionbutton">Open</button></Link>
                  </div>
               </div>
            </div>
         </div>
    );
  }
}

export default App
