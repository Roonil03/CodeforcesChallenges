```cpp
#include <bits/stdc++.h>

using namespace std;

#pragma GCC optimize ("O3")
// #pragma GCC target("avx2, bmi, bmi2, popcnt")

#define sz(x) ((int)x.size())
#define ll long long
#define pii pair<int, int>
#define pll pair<ll, ll>
#define vi vector<int>
#define vll vector<long long>

template<typename T>
istream& operator>>(istream& is, vector<T>& v){
    for (T& x : v) is >> x;
    return is;
}

template<typename T>
ostream& operator<<(ostream& os, const vector<T>& v){
    for (int i = 0; i < v.size(); i++){
        os << v[i];
        if (i + 1 < v.size()) os << " ";
    }
    return os;
}

const int Z = 10000;
const int INF = 0x3f3f3f3f;

int x_, y_;

void precomp(){
    int a, b, c, s, s_;
    x_ = 0, y_ = 0, s_ = INF;
    for (a = 0; a <= Z; a++) {
        for (b = 0; a + b + a <= Z; b++) {
            c = Z - a - b;
            s = max(max(a, b / 2), c);
            s = max(s, min(b, a * 2));
            s = max(s, min(b / 2, c * 2));
            if (s_ > s) {
                s_ = s;
                x_ = a;
                y_ = a + b;
            }
        }
    }
}

int query(int x, int y, int& k) {
    int d;
    cout << x << " " << y << "\n";
    cout.flush();
    k++;
    cin >> d;
    return d;
}

void solve(){
    int k = 0;
    int d = query(x_, y_, k);
    int x, y;
    if (d == 0) {
        x = x_;
        y = y_;
    } else if (x_ + d > y_ - d) {
        x = x_ - d;
        y = y_ + d;
    } else if (y_ + d > Z) {
        x = x_ + d;
        y = y_ - d;
    } else if (x_ - d < 0) {
        if ((y_ - d) - (x_ + d) > (y_ + d) - (y_ - d)) {
            int d_ = query(x_ + d, y_, k);
            if (d_ == 0) {
                x = x_ + d;
                y = x_ + d;
            } else {
                x = y_ - d;
                y = y_ + d;
            }
        } else {
            int d_ = query((x_ + y_) / 2, y_ + d, k);
            if (d_ == 0) {
                x = (x_ + y_) / 2;
                y = y_ + d;
            } else {
                x = x_ + d;
                y = y_ - d;
            }
        }
    } else {
        int d_ = query(x_ + 1, y_ - 1, k);
        if (d_ < d) {
            x = x_ + d;
            y = y_ - d;
        } else {
            x = x_ - d;
            y = y_ + d;
        }
    }
    while (k < 10) {
        query(x, y, k);
    }
}

int main(){
    ios::sync_with_stdio(0);
    cin.tie(0);
    cout.tie(nullptr);
    cin.exceptions(cin.failbit);
    int t;
    cin >> t;
    precomp();
    while(t--){
        solve();
    }
    return 0;
}
```
```cpp
// Solution referrenced from https://codeforces.com/profile/rainboy

/*
This algorithm plays an interactive guessing game on a 1D grid to pinpoint a hidden point $(x, y)$. It uses an initial static 
configuration $(x\_, y\_)$ precomputed during `precomp` using brute-force search over the constrained grid up to $Z = 10000$. 
This optimal initial query serves to strictly bound the worst-case maximum size of the search partitions. Once the interactive 
solving phase starts, it submits $(x\_, y\_)$ to receive a distance feedback $d$. Using this distance and predefined bounds 
matching the partition metrics, the algorithm deduces the actual $(x, y)$ coordinate. By testing targeted geometric offsets 
depending on the distance and boundary restrictions (e.g. `x_ - d < 0` or `y_ + d > Z`), the subsequent response $d'$ enables 
discrimination between the possible target offsets without fail. Finally, it uses up the remaining operations allowance of 10 
queries by continually querying the determined correct location. The overall time complexity per testcase is $O(1)$ query 
interactions taking initially $O(Z^2)$ space and time during the static `precomp` phase. The space complexity during testcase 
execution is strictly $O(1)$.
*/
```