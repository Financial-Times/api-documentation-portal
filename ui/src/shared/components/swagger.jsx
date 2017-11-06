import React from 'react';

import { SwaggerUIBundle as SwaggerUi, SwaggerUIStandalonePreset } from 'swagger-ui-dist'

export default class SwaggerUI extends React.Component {

  componentDidMount() {
    SwaggerUi({
      dom_id: '#swaggerContainer',
      url: this.props.url,
      spec: undefined,
      presets: [
        SwaggerUi.presets.apis,
        SwaggerUIStandalonePreset
      ],
      plugins: [
        SwaggerUi.plugins.DownloadUrl
      ],
      layout: 'StandaloneLayout',
      validatorUrl: null
    });
  }

  render() {
    return (
         <div id="swaggerContainer" />
    );
  }
}
