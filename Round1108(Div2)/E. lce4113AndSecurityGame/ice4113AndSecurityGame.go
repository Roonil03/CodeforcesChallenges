package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var (
	in  = NewFastScanner()
	out = bufio.NewWriter(os.Stdout)
)

type FastScanner struct {
	r *bufio.Reader
}

func NewFastScanner() *FastScanner {
	return &FastScanner{r: bufio.NewReaderSize(os.Stdin, 1<<20)}
}

func (fs *FastScanner) rInt() int {
	sign, val := 1, 0
	c, err := fs.r.ReadByte()
	for (c < '0' || c > '9') && c != '-' {
		if err != nil {
			return 0
		}
		c, err = fs.r.ReadByte()
	}
	if c == '-' {
		sign = -1
		c, _ = fs.r.ReadByte()
	}
	for c >= '0' && c <= '9' {
		val = val*10 + int(c-'0')
		c, err = fs.r.ReadByte()
		if err != nil {
			break
		}
	}
	return sign * val
}

func (fs *FastScanner) rInt64() int64 {
	sign, val := int64(1), int64(0)
	c, err := fs.r.ReadByte()
	for (c < '0' || c > '9') && c != '-' {
		if err != nil {
			return 0
		}
		c, err = fs.r.ReadByte()
	}
	if c == '-' {
		sign = -1
		c, _ = fs.r.ReadByte()
	}
	for c >= '0' && c <= '9' {
		val = val*10 + int64(c-'0')
		c, err = fs.r.ReadByte()
		if err != nil {
			break
		}
	}
	return sign * val
}

func precomp() {}

func solve() {
	b := int64(1)
	fmt.Fprintln(out, b)
	out.Flush()
	x := in.rInt64()
	if x == -1 {
		os.Exit(0)
	}
	if x^b != 0 {
		d := x ^ b
		d &= -d
		fmt.Fprintln(out, 0, d)
		out.Flush()
		y := in.rInt64()
		if y == -1 {
			os.Exit(0)
		}
		if (y & d) == (x & d) {
			fmt.Fprintln(out, 0)
		} else {
			fmt.Fprintln(out, 1)
		}
		out.Flush()
	} else {
		g := rand.Int63n(1<<30) | 1
		if g == 1 {
			g = 3
		}
		fmt.Fprintln(out, 0, g)
		out.Flush()
		y := in.rInt64()
		if y == -1 {
			os.Exit(0)
		}
		if y-(y&b) == 0 {
			fmt.Fprintln(out, 0)
		} else if (y^g)-((y^g)&b) == 0 {
			fmt.Fprintln(out, 1)
		} else {
			fmt.Fprintln(out, 1-(y&b))
		}
		out.Flush()
	}
}

func main() {
	out = bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()
	precomp()
	t := in.rInt()
	for t > 0 {
		t--
		solve()
	}
}

// Solution not working in go, going into Idleness errors. Check solution by: https://codeforces.com/profile/Nachia

/*
#ifdef NACHIA
#define _GLIBCXX_DEBUG
#else
// disable assert
#define NDEBUG
#endif
#include <iostream>
#include <string>
#include <vector>
#include <algorithm>
using namespace std;
using ll = long long;
const ll INF = 1ll << 60;
#define REP(i,n) for(ll i=0; i<ll(n); i++)
template <class T> using V = vector<T>;
template <class A, class B> void chmax(A& l, const B& r){ if(l < r) l = r; }
template <class A, class B> void chmin(A& l, const B& r){ if(r < l) l = r; }

#include <unordered_map>
#include <cassert>
#include <cstdint>
#include <array>


namespace nachia{

int Popcount(unsigned long long c){
#ifdef __GNUC__
    return __builtin_popcountll(c);
#else
    c = (c & (~0ull/3)) + ((c >> 1) & (~0ull/3));
    c = (c & (~0ull/5)) + ((c >> 2) & (~0ull/5));
    c = (c & (~0ull/17)) + ((c >> 4) & (~0ull/17));
    c = (c * (~0ull/257)) >> 56;
    return c;
#endif
}

// please ensure x != 0
int MsbIndex(unsigned long long x){
    assert(x != 0ull);
#ifdef __GNUC__
    return 63 - __builtin_clzll(x);
#else
    using u64 = unsigned long long;
    int q = (x >> 32) ? 32 : 0;
    auto m = x >> q;
    constexpr u64 hi = 0x88888888;
    constexpr u64 mi = 0x11111111;
    m = (((m | ~(hi - (m & ~hi))) & hi) * mi) >> 35;
    m = (((m | ~(hi - (m & ~hi))) & hi) * mi) >> 31;
    q += (m & 0xf) << 2;
    q += 0x3333333322221100 >> (((x >> q) & 0xf) << 2) & 0xf;
    return q;
#endif
}

// please ensure x != 0
int LsbIndex(unsigned long long x){
    assert(x != 0ull);
#ifdef __GNUC__
    return __builtin_ctzll(x);
#else
    return MsbIndex(x & -x);
#endif
}

}


namespace nachia{

class Xoshiro256pp{
public:

    using i32 = int32_t;
    using u32 = uint32_t;
    using i64 = int64_t;
    using u64 = uint64_t;


private:
    std::array<u64, 4> s;

    // https://prng.di.unimi.it/xoshiro256plusplus.c
    static inline uint64_t rotl(const uint64_t x, int k) noexcept {
        return (x << k) | (x >> (64 - k));
    }
    inline uint64_t gen(void) noexcept {
        const uint64_t result = rotl(s[0] + s[3], 23) + s[0];
        const uint64_t t = s[1] << 17;
        s[2] ^= s[0];
        s[3] ^= s[1];
        s[1] ^= s[2];
        s[0] ^= s[3];
        s[2] ^= t;
        s[3] = rotl(s[3], 45);
        return result;
    }

    // https://xoshiro.di.unimi.it/splitmix64.c
    u64 splitmix64(u64& x) {
        u64 z = (x += 0x9e3779b97f4a7c15);
        z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9;
        z = (z ^ (z >> 27)) * 0x94d049bb133111eb;
        return z ^ (z >> 31);
    }

    // generate x : 0 <= x <= r
    u64 random_unsigned_from_zero(u64 r){
        if(!r) return 0;
        u64 mask = 1ull << MsbIndex(r);
        mask += mask - 1;
        while(true){
            auto res = rng64() & mask;
            if(res <= r) return res;
        }
    }

public:

    void seed(u64 x = 7001){
        assert(x != 0);
        s[0] = x;
        for(int i=1; i<4; i++) s[i] = splitmix64(x);
    }

    std::array<u64, 4> getState() const { return s; }
    void setState(std::array<u64, 4> a){ s = a; }

    Xoshiro256pp(){ seed(); }

    u64 rng64() { return gen(); }
    u64 operator()(){ return gen(); }

    // generate x : l <= x <= r
    u64 random_unsigned(u64 l,u64 r){
        assert(l<=r);
        return l + random_unsigned_from_zero(r - l);
    }

    // generate x : l <= x <= r
    i64 random_signed(i64 l,i64 r){
        assert(l<=r);
        return (i64)(random_unsigned_from_zero((u64)r - (u64)l) + (u64)l);
    }


    // permute x : n_left <= x <= n_right
    // output r from the front
    template<class Int>
    std::vector<Int> random_nPr(Int n_left, Int n_right, Int r){
        Int n = n_right-n_left;

        assert(n>=0);
        assert(r<=(1ll<<27));
        if(r==0) return {};
        assert(n>=r-1);

        std::vector<Int> V;
        std::unordered_map<Int,Int> G;
        for(int i=0; i<r; i++){
            Int p = random_signed(i,n);
            Int x = p - G[p];
            V.push_back(x);
            G[p] = p - (i - G[i]);
        }

        for(Int& v : V) v+=n_left;
        return V;
    }

    // shuffle using swaps
    template<class E>
    void shuffle_inplace(std::vector<E>& V){
        size_t N = V.size();
        for(size_t i=1; i<N; i++){
            std::swap(V[i], V[random_unsigned(0, i)]);
        }
    }

    template<class E>
    std::vector<E> shuffled(const std::vector<E>& V){
        auto cp = V;
        shuffle_inplace(cp);
        return cp;
    }

};

} // namespace nachia
#include <random>
#include <chrono>
nachia::Xoshiro256pp rng;
namespace RngInitInstance { struct RngInit { RngInit(){
    unsigned long long seed1 = std::random_device()();
    unsigned long long seed2 = std::chrono::high_resolution_clock::now().time_since_epoch().count();
    auto s = seed1 ^ seed2;
    #ifdef NACHIA
    std::cerr << "seed = " << s << std::endl;
    #endif
    rng.seed(s);
} } a; }

void testcase(){
  ll b = 1;
  cout << b << endl;
  ll x; cin >> x;
  if(x ^ b){
    ll d = x ^ b;
    d &= -d;
    cout << 0 << " " << d << endl;
    ll y; cin >> y;
    if((y & d) == (x & d)){
      cout << 0 << endl;
    } else {
      cout << 1 << endl;
    }
  }
  else {
    ll g = (rng() % (1ll << 30)) | 1;
    cout << 0 << " " << g << endl;
    ll y; cin >> y;
    if((y - (y & b)) == 0){
      cout << 0 << endl;
    } else if(((y ^ g) - ((y ^ g) & b)) == 0){
      cout << 1 << endl;
    } else {
      cout << (1 - (y & b)) << endl;
    }
  }
}

int main(){
  cin.tie(0)->sync_with_stdio(0);
  ll T; cin >> T; REP(t,T)
  testcase();
  return 0;
}
*/

/*
This interactive algorithm determines a hidden state by initially querying with a baseline mask $b = 1$. If the response
reveals a discrepancy ($x \oplus b \neq 0$), it precisely isolates the least significant differing bit, $d$, and leverages
it in a follow-up query to conclusively distinguish the layout based on target subset parity. Conversely, if the initial
query results in a perfect match, the algorithm assumes ambiguity exists and deploys a full-mask query $g$ (configured to
guarantee a set LSB) to confidently extract the remaining structural bits without overlap errors. The time complexity per
testcase is $O(1)$ since determining the bits relies purely on constant-time logical and arithmetic operations without iterating.
The space complexity is rigorously $O(1)$ as the execution requires solely a handful of primitive integer variables for
query generation and response tracking.
*/
