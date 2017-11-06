import React from "react";
import {observer} from "mobx-react";

/**
 *
 * @type {Observer}
 */
@observer
class App extends React.Component {

  render(){
    let app = this.props.app;

    return (
      <section className="app">
         <h1 id={app.name}>{app.title}</h1>
         {app.tags.sort().map(tag => {
           return (
             <span key={tag} className="tag">{tag}</span>
           )
         })}
         <div className="u-cf"></div>
         <p>{app.description}</p>
         <a href={"documentation.html#/docs/" + app.name}><button>API Documentation 📝</button></a>
      </section>
    );
  }
}

export default App
