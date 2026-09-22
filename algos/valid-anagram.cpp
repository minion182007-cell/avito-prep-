class Solution {
public:
    bool isAnagram(string s, string t) {
        if (s.length() != t.length()) return false;
        unordered_map<char,int> seen;
        unordered_map<char,int> seen1;
        for (int i=0;i<s.length();i++){
            seen[s[i]]++;
        } 
        for (int i=0;i<t.length();i++){
            seen1[t[i]]++;
        }
        if (seen1==seen){
            return true;
        }
        return false;
    }
};
