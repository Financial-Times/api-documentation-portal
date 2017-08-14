import React from "react";

export default () => {
  return (
      <header className="o-header-services"
        data-o-component="o-header"
      >
         <div className="o-header-services__top o-header-services__container">
            <div className="o-header-services__ftlogo" />
            <div className="o-header-services__title">
               <h2 className="o-header-services__product-name"><a href="/">UPP API Hub</a></h2>
            </div>
            <div className="o-header-services__related-content">
               {/*<a className="o-header-services__related-content-link" href="#">Related site</a>
               <a className="o-header-services__related-content-link" href="#">Sign in</a>*/}
            </div>
         </div>
      </header>
  );
};
