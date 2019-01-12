type config = {
  apiEndpoint: string;
};

const config: config = {
  apiEndpoint: (process.env.NODE_ENV === 'production') ? 'http://todone.brennonloveless.com:9090' : 'http://localhost:9090',
};

export default config;
