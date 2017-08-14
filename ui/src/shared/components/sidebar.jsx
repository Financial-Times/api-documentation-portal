import React from "react";
import {observer} from "mobx-react";

export default ({apps}) => {
  return (
      <div className="o-techdocs-sidebar">
         <ul className="o-techdocs-nav">
            {/* {apps.map((app) => {
               var href = "#" + app.name
               return (<li key={app.title}><a href={href}>{app.title}</a></li>)
            })} */}
         </ul>
      </div>
  );
};
