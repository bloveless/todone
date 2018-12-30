const HtmlWebpackPlugin = require('html-webpack-plugin');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const path = require('path');

const mode = (process.env.NODE_ENV === 'production') ? 'production' : 'development';

module.exports = {
    mode,

    entry: './src/index.tsx',

    devServer: {
        host: '0.0.0.0',
        port: 8080,
        proxy: {
            '/**': {  //catch all requests
                target: '/index.html',  //default target
                secure: false,
                bypass: function (req, res, opt) {
                    //your custom code to check for any exceptions
                    //console.log('bypass check', {req: req, res:res, opt: opt});
                    if (req.path.indexOf('/img/') !== -1 || req.path.indexOf('/public/') !== -1) {
                        return '/'
                    }

                    if (req.headers.accept.indexOf('html') !== -1) {
                        return '/index.html';
                    }
                },
            },
        },
    },

    output: {
        publicPath: '/',
        filename: (mode === 'development') ? '[name].js' : '[name].[contentHash].js',
        path: path.resolve(__dirname, 'dist'),
    },

    // Enable sourcemaps for debugging webpack's output.
    devtool: 'source-map',

    resolve: {
        // Add '.ts' and '.tsx' as resolvable extensions.
        extensions: ['.ts', '.tsx', '.js', '.json'],
    },

    module: {
        rules: [
            // All files with a '.ts' or '.tsx' extension will be handled by 'awesome-typescript-loader'.
            {
                test: /\.tsx?$/,
                loader: 'awesome-typescript-loader',
            },

            // All output '.js' files will have any sourcemaps re-processed by 'source-map-loader'.
            {
                enforce: 'pre',
                test: /\.js$/,
                loader: 'source-map-loader',
            },


            // Compile and include scss files.
            {
                test: /\.scss$/,
                use: [
                    'css-hot-loader',
                    MiniCssExtractPlugin.loader,
                    {
                        loader: 'css-loader',
                        options: {
                            importLoaders: 2,
                            sourceMap: true,
                        },
                    },
                    {
                        loader: 'postcss-loader',
                        options: {
                            plugins: () => [
                                require('autoprefixer'),
                            ],
                            sourceMap: true,
                        },
                    },
                    {
                        loader: 'sass-loader',
                        options: {
                            sourceMap: true,
                        },
                    },
                ],
            },
        ],
    },

    plugins: [
        new CleanWebpackPlugin(path.resolve(__dirname, 'dist')),
        new HtmlWebpackPlugin({
            template: 'src/index.html'
        }),
        new MiniCssExtractPlugin({
            filename: `styles/[name].css`,
        }),
    ]
};
