type config = {
  apiEndpoint: string;
};

const config: config = {
  apiEndpoint: (process.env.NODE_ENV === 'production') ? 'http://todone-api.brennonloveless.com:9090/v1' : 'http://localhost:9090/v1',
};

export default config;
