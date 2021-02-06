

// var s = require("./D3628302AB6B1EDFB504EB051053A6B3.js"),
//     a = 0,
//     n = function() {
//     var e = (Date.parse(new Date()) + a) / 1e3 + "";
//     console.log(e);
//     return s("zfsw_" + e.substring(0, e.length - 1));
// };

function zftsl(){
    // var s = require("./D3628302AB6B1EDFB504EB051053A6B3.js")
    var a= 0
    var e = (Date.parse(new Date()) + a) / 1e3 + "";
    return s("zfsw_" + e.substring(0, e.length - 1));
}
function s(n){
    return l(v(n))
}
function r(n, e) {
    var r = (65535 & n) + (65535 & e);
    return (n >> 16) + (e >> 16) + (r >> 16) << 16 | 65535 & r;
}
function t(n, e, t, u, o, f) {
    return r((i = r(r(e, n), r(u, f))) << (c = o) | i >>> 32 - c, t);
    var i, c;
}
function u(n, e, r, u, o, f, i) {
    return t(e & r | ~e & u, n, e, o, f, i);
}
function o(n, e, r, u, o, f, i) {
    return t(e & u | r & ~u, n, e, o, f, i);
}
function f(n, e, r, u, o, f, i) {
    return t(e ^ r ^ u, n, e, o, f, i);
}
function i(n, e, r, u, o, f, i) {
    return t(r ^ (e | ~u), n, e, o, f, i);
}
function c(n, e) {
    var t, c, a, d;
    n[e >> 5] |= 128 << e % 32, n[14 + (e + 64 >>> 9 << 4)] = e;
    for (var l = 1732584193, h = -271733879, v = -1732584194, g = 271733878, m = 0; m < n.length; m += 16) {
        l = u(t = l, c = h, a = v, d = g, n[m], 7, -680876936), g = u(g, l, h, v, n[m + 1], 12, -389564586), v = u(v, g, l, h, n[m + 2], 17, 606105819), h = u(h, v, g, l, n[m + 3], 22, -1044525330), l = u(l, h, v, g, n[m + 4], 7, -176418897), g = u(g, l, h, v, n[m + 5], 12, 1200080426), v = u(v, g, l, h, n[m + 6], 17, -1473231341), h = u(h, v, g, l, n[m + 7], 22, -45705983), l = u(l, h, v, g, n[m + 8], 7, 1770035416), g = u(g, l, h, v, n[m + 9], 12, -1958414417), v = u(v, g, l, h, n[m + 10], 17, -42063), h = u(h, v, g, l, n[m + 11], 22, -1990404162), l = u(l, h, v, g, n[m + 12], 7, 1804603682), g = u(g, l, h, v, n[m + 13], 12, -40341101), v = u(v, g, l, h, n[m + 14], 17, -1502002290), l = o(l, h = u(h, v, g, l, n[m + 15], 22, 1236535329), v, g, n[m + 1], 5, -165796510), g = o(g, l, h, v, n[m + 6], 9, -1069501632), v = o(v, g, l, h, n[m + 11], 14, 643717713), h = o(h, v, g, l, n[m], 20, -373897302), l = o(l, h, v, g, n[m + 5], 5, -701558691), g = o(g, l, h, v, n[m + 10], 9, 38016083), v = o(v, g, l, h, n[m + 15], 14, -660478335), h = o(h, v, g, l, n[m + 4], 20, -405537848), l = o(l, h, v, g, n[m + 9], 5, 568446438), g = o(g, l, h, v, n[m + 14], 9, -1019803690), v = o(v, g, l, h, n[m + 3], 14, -187363961), h = o(h, v, g, l, n[m + 8], 20, 1163531501), l = o(l, h, v, g, n[m + 13], 5, -1444681467), g = o(g, l, h, v, n[m + 2], 9, -51403784), v = o(v, g, l, h, n[m + 7], 14, 1735328473), l = f(l, h = o(h, v, g, l, n[m + 12], 20, -1926607734), v, g, n[m + 5], 4, -378558), g = f(g, l, h, v, n[m + 8], 11, -2022574463), v = f(v, g, l, h, n[m + 11], 16, 1839030562), h = f(h, v, g, l, n[m + 14], 23, -35309556), l = f(l, h, v, g, n[m + 1], 4, -1530992060), g = f(g, l, h, v, n[m + 4], 11, 1272893353), v = f(v, g, l, h, n[m + 7], 16, -155497632), h = f(h, v, g, l, n[m + 10], 23, -1094730640), l = f(l, h, v, g, n[m + 13], 4, 681279174), g = f(g, l, h, v, n[m], 11, -358537222), v = f(v, g, l, h, n[m + 3], 16, -722521979), h = f(h, v, g, l, n[m + 6], 23, 76029189), l = f(l, h, v, g, n[m + 9], 4, -640364487), g = f(g, l, h, v, n[m + 12], 11, -421815835), v = f(v, g, l, h, n[m + 15], 16, 530742520), l = i(l, h = f(h, v, g, l, n[m + 2], 23, -995338651), v, g, n[m], 6, -198630844), g = i(g, l, h, v, n[m + 7], 10, 1126891415), v = i(v, g, l, h, n[m + 14], 15, -1416354905), h = i(h, v, g, l, n[m + 5], 21, -57434055), l = i(l, h, v, g, n[m + 12], 6, 1700485571), g = i(g, l, h, v, n[m + 3], 10, -1894986606), v = i(v, g, l, h, n[m + 10], 15, -1051523), h = i(h, v, g, l, n[m + 1], 21, -2054922799), l = i(l, h, v, g, n[m + 8], 6, 1873313359), g = i(g, l, h, v, n[m + 15], 10, -30611744), v = i(v, g, l, h, n[m + 6], 15, -1560198380), h = i(h, v, g, l, n[m + 13], 21, 1309151649), l = i(l, h, v, g, n[m + 4], 6, -145523070), g = i(g, l, h, v, n[m + 11], 10, -1120210379), v = i(v, g, l, h, n[m + 2], 15, 718787259), h = i(h, v, g, l, n[m + 9], 21, -343485551), l = r(l, t), h = r(h, c), v = r(v, a), g = r(g, d);
    }return [l, h, v, g];
}
function a(n) {
    for (var e = "", r = 32 * n.length, t = 0; t < r; t += 8) {
        e += String.fromCharCode(n[t >> 5] >>> t % 32 & 255);
    }return e;
}
function d(n) {
    var e = [];
    for (e[(n.length >> 2) - 1] = void 0, t = 0; t < e.length; t += 1) {
        e[t] = 0;
    }for (var r = 8 * n.length, t = 0; t < r; t += 8) {
        e[t >> 5] |= (255 & n.charCodeAt(t / 8)) << t % 32;
    }return e;
}
function l(n) {
    for (var e, r = "0123456789abcdef", t = "", u = 0; u < n.length; u += 1) {
        e = n.charCodeAt(u), t += r.charAt(e >>> 4 & 15) + r.charAt(15 & e);
    }return t;
}
function h(n) {
    return unescape(encodeURIComponent(n));
}
function v(n) {
    return a(c(d(e = h(n)), 8 * e.length));
    var e;
}
function g(n, e) {
    return function (n, e) {
        var r,
            t,
            u = d(n),
            o = [],
            f = [];
        for (o[15] = f[15] = void 0, 16 < u.length && (u = c(u, 8 * n.length)), r = 0; r < 16; r += 1) {
            o[r] = 909522486 ^ u[r], f[r] = 1549556828 ^ u[r];
        }return t = c(o.concat(d(e)), 512 + 8 * e.length), a(c(f.concat(t), 640));
    }(h(n), h(e));
}
function m(n, e, r) {
    return e ? r ? g(e, n) : l(g(e, n)) : r ? v(n) : l(v(n));
}
// console.log(zftsl())
// console.log(n())
