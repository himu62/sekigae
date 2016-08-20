var webpack = require('webpack');
var path = require('path');
var ExtractTextPlugin = require('extract-text-webpack-plugin');

var plugins = [];
var css_loader = 'css-loader!sass-loader';
if(process.env.NODE_ENV === 'production') {
	plugins = [
		new webpack.optimize.UglifyJsPlugin(),
		new webpack.optimize.DedupePlugin(),
		new webpack.DefinePlugin({
			'process.env': {
				'NODE_ENV': JSON.stringify('production')
			}
		})
	];
	css_loader = 'css-loader?minimize!sass-loader';
}

module.exports = [
	{
		entry: path.join(__dirname, 'client/index.tsx'),
		output: {
			filename: 'public/js/bundle.js'
		},
		resolve: {
			extensions: ['', '.tsx', '.ts', '.js']
		},
		module: {
			loaders: [
				{
					test: /\.tsx?$/,
					loader: 'ts-loader',
					exclude: /node_modules/
				}
			]
		},
		plugins: plugins
	},
	{
		entry: path.join(__dirname, 'style/style.scss'),
		output: {
			filename: 'public/css/style.css'
		},
		resolve: {
			extensions: ['', '.scss', '.css']
		},
		module: {
			loaders: [
				{
					test: /\.scss$/,
					loader: ExtractTextPlugin.extract('style-loader', css_loader)
				}
			]
		},
		plugins: [
			new ExtractTextPlugin('public/css/style.css')
		]
	}
];
