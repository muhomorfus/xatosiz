import * as openapi from '@/openapi';

export default new openapi.DefaultApi(new openapi.Configuration({
    basePath: 'https://api.xatosiz.lownie.su',
    // basePath: 'http://localhost:3010',
}));